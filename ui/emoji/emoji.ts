interface EmojiInfo {
    /**
     * emoji 符号
     */
    emoji: string;
    /**
     * 编码
     */
    code: string,
    /**
     * 名称
     */
    name: string,
    /**
     * 描述
     */
    description: string,
    /**
     * Unicode 编码
     */
    unicode: number | number[];
    /**
     * 语义化版本
     * major：不兼容的 API 修改
     * minor：向下兼容的功能性新增
     * patch：向下兼容的问题修正
     */
    semver?: 'major' | 'minor' | 'patch';
}

enum Sev {
    MAJOR = 'MAJOR',
    MINOR = 'MINOR',
    PATCH = 'PATCH'
}