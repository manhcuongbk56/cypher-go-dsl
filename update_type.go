package cypher

type UpdateType string

const (
	UPDATE_TYPE_DELETE        UpdateType = "delete"
	UPDATE_TYPE_DETACH_DELETE            = "detachDelete"
	UPDATE_TYPE_SET                      = "OperationSet"
	UPDATE_TYPE_REMOVE                   = "OperationRemove"
	UPDATE_TYPE_CREATE                   = "create"
	UPDATE_TYPE_MERGE                    = "merge"
)
