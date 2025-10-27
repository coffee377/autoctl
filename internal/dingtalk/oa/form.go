package oa

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	dingtalkworkflow10 "github.com/alibabacloud-go/dingtalk/workflow_1_0"
)

// FormComponent 钉钉表单组件结构（与钉钉返回的表单组件字段对应）
// 假设钉钉返回的组件结构包含 Id 和 Value（指针类型，因为可能为nil）
type FormComponent struct {
	Id    *string // 表单组件ID（如"TextField_1"）
	Value *string // 表单值（字符串类型，需要转换）
}

// FieldMapper 字段映射规则：描述表单组件到实体字段的映射关系
type FieldMapper struct {
	ComponentId string         // 钉钉表单组件ID（如"TextField_1"）
	FieldName   string         // 实体对象的字段名（如"ProjectName"）
	Converter   ValueConverter // 类型转换函数（将字符串转为字段类型）
	Pointer     bool
}

// ValueConverter 类型转换接口：将表单字符串值转换为目标类型
type ValueConverter func(raw string, pointer bool) (interface{}, error)

// StringConverter 字符串转换器（默认，直接trim空格）
func StringConverter(raw string, pointer bool) (interface{}, error) {
	trim := strings.Trim(raw, " ")
	if pointer {
		return &trim, nil
	}
	return trim, nil
}

// DateConverter 时间转换器（支持"2006-01-02"格式，返回*time.Time）
func DateConverter(layout string, loc *time.Location) ValueConverter {
	return func(raw string, pointer bool) (interface{}, error) {
		raw = strings.Trim(raw, " ")
		if raw == "" {
			return nil, nil // 空值返回nil（适配指针类型）
		}
		t, err := time.ParseInLocation(layout, raw, loc)
		if err != nil {
			return nil, fmt.Errorf("parse date failed: %w", err)
		}
		if pointer {
			return &t, nil
		}
		return t, nil
	}
}

// Float64Converter 浮点数转换器（返回float64）
func Float64Converter(raw string, pointer bool) (interface{}, error) {
	raw = strings.Trim(raw, " ")
	if raw == "" {
		r := 0.0
		if pointer {
			return &r, nil
		}
		return r, nil
	}
	val, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return nil, fmt.Errorf("parse float failed: %w", err)
	}
	if pointer {
		return &val, nil
	}
	return val, nil
}

// MapFormToEntity 将钉钉表单组件列表映射到实体对象
// res：钉钉返回的数据
// mappers：字段映射规则
// entity：目标实体指针（必须是指针类型，否则无法赋值）
func MapFormToEntity(res *dingtalkworkflow10.GetProcessInstanceResponseBodyResult, mappers []FieldMapper, entity interface{}) error {
	components := res.FormComponentValues
	// 检查实体是否为指针
	entityVal := reflect.ValueOf(entity)
	if entityVal.Kind() != reflect.Ptr || entityVal.IsNil() {
		return errors.New("entity must be a non-nil pointer")
	}
	// 获取实体的实际类型（解引用指针）
	entityElem := entityVal.Elem()
	if entityElem.Kind() != reflect.Struct {
		return errors.New("entity must be a struct pointer")
	}

	// 将组件列表转为map（ComponentId -> Value），方便查询
	componentMap := make(map[string]string)
	for _, comp := range components {
		if comp.Id != nil && comp.Value != nil {
			// todo comp.ExtValue 获取连接器绑定的数据
			componentMap[*comp.Id] = *comp.Value
		}
	}

	// 遍历映射规则，填充字段
	for _, mapper := range mappers {
		// 1. 查找表单组件值
		rawValue, exists := componentMap[mapper.ComponentId]
		if !exists {
			continue // 组件不存在，跳过（或根据需求报错）
		}

		// 2. 转换值为目标类型
		convertedValue, err := mapper.Converter(rawValue, mapper.Pointer)

		if err != nil {
			return fmt.Errorf("field %s convert failed: %w", mapper.FieldName, err)
		}

		// 3. 反射设置实体字段
		field := entityElem.FieldByName(mapper.FieldName)
		if !field.IsValid() {
			return fmt.Errorf("entity has no field: %s", mapper.FieldName)
		}
		if !field.CanSet() {
			return fmt.Errorf("field %s is unexported (must start with uppercase)", mapper.FieldName)
		}

		// 处理值类型（如convertedValue是*time.Time，字段是*time.Time）
		val := reflect.ValueOf(convertedValue)
		if val.Type() != field.Type() {
			return fmt.Errorf("field %s type mismatch: expect %s, got %s",
				mapper.FieldName, field.Type(), val.Type())
		}
		field.Set(val)
	}

	return nil
}
