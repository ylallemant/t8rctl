package api

var (
	TAG_DEFAULT_VALUE = "none"
	TAG_CLUSTER_GROUP = "cluster_group"
	TAG_CLUSTER_ID    = "cluster_id"
	TAG_CLUSTER_STAGE = "cluster_stage"
	TAG_STAGE         = "stage"
	TAG_CELL          = "cell"
	TAG_CELL_ID       = "solution_id"
	TAG_CELL_TENANT   = "solution_tenant"
	TAG_CELL_REGION   = "solution_region"
	TAG_DATATIER      = "datatier"
)

func CastString(raw *string) string {
	value := *raw

	if value == "" {
		return TAG_DEFAULT_VALUE
	}

	return value
}

func ReadTag(tags map[string]string, tag string) string {
	val, ok := tags[tag]
	if ok {
		return val
	}

	return TAG_DEFAULT_VALUE
}
