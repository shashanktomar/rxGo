package fuseable

type SimpleQueue interface{
  offerValue(value interface{}) bool
  offerValues(v1 interface{}, v2 interface{}) bool
  poll() (interface{}, error)
  isEmpty() bool
  clear()
}
