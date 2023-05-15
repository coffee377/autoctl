package lib

// Blob stored in git.
type Blob interface {
	// ID of the blob.
	ID() string
	// Name of the blob.
	Name() string
	// Path (relative) to the blob.
	Path() string
	//String returns a printable id.
	String() string
}

// Commit in the repo.
// All commit based APIs in Repo interface accepts this interface.
// It gives the implementations the ability to optimise the access to commits.
// For example, a libgit2 based Repo implementation we use by default caches the access to commit tree.
type Commit interface {
	ID() string
	String() string
}

// Reference to a tree in the repository.
// For example, if you consider a git repository
// this could be pointing to a branch, tag or commit.
type Reference interface {
	Name() string
	SymbolicName() string
}

// DiffDelta is a single delta in a git diff.
type DiffDelta struct {
	// NewFile path of the delta
	NewFile string
	// OldFile path of the delta
	OldFile string
}

// BlobWalkCallback used for discovering blobs in a commit tree.
type BlobWalkCallback func(Blob) error

// Repo defines the set of interactions with the git repository.
type Repo interface {
	// GetCommit returns the commit object for the specified SHA.
	GetCommit(sha string) (Commit, error)
	// Path of the repository.
	Path() string
	// Diff gets the diff between two commits.
	Diff(a, b Commit) ([]*DiffDelta, error)
	// DiffMergeBase gets the diff between the merge base of from and to and, to.
	// In other words, diff contains the deltas of changes occurred in 'to' commit tree
	// since it diverged from 'from' commit tree.
	DiffMergeBase(from, to Commit) ([]*DiffDelta, error)
	// DiffWorkspace gets the changes in current workspace.
	// This should include untracked changes.
	DiffWorkspace() ([]*DiffDelta, error)
	// Changes returns an array of DiffDelta objects representing the changes
	// in the specified commit.
	// Return an empty array if the specified commit is the first commit
	// in the repo.
	Changes(c Commit) ([]*DiffDelta, error)
	// WalkBlobs invokes the callback for each blob reachable from the commit tree.
	WalkBlobs(a Commit, callback BlobWalkCallback) error
	// BlobContents of specified blob.
	BlobContents(blob Blob) ([]byte, error)
	// BlobContentsByPath gets the blob contents from a specific git tree.
	BlobContentsFromTree(commit Commit, path string) ([]byte, error)
	// EntryID of a git object in path.
	// ID is resolved from the commit tree of the specified commit.
	EntryID(commit Commit, path string) (string, error)
	// BranchCommit returns the last commit for the specified branch.
	BranchCommit(name string) (Commit, error)
	// CurrentBranch returns the name of current branch.
	CurrentBranch() (string, error)
	// CurrentBranchCommit returns the last commit for the current branch.
	CurrentBranchCommit() (Commit, error)
	// IsEmpty informs if the current repository is empty or not.
	IsEmpty() (bool, error)
	// FindAllFilesInWorkspace returns all files in repository matching given pathSpec, including untracked files.
	FindAllFilesInWorkspace(pathSpec []string) ([]string, error)
	// EnsureSafeWorkspace returns an error workspace is in a safe state
	// for operations requiring a checkout.
	// For example, in git repositories we consider uncommitted changes or
	// a detached head is an unsafe state.
	EnsureSafeWorkspace() error
	// Checkout specified commit into workspace.
	// Also returns a reference to the previous tree pointed by current workspace.
	Checkout(commit Commit) (Reference, error)
	// CheckoutReference checks out the specified reference into workspace.
	CheckoutReference(Reference) error
	// MergeBase returns the merge base of two commits.
	MergeBase(a, b Commit) (Commit, error)
}

// Module represents a single module in the repository.
type Module struct {
	//metadata   *moduleMetadata
	version    string
	requires   Modules
	requiredBy Modules
}

// Modules is an array of Module.
type Modules []*Module
