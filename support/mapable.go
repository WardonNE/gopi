package support

type Mapable interface {
	ToMap() map[any]any
}
