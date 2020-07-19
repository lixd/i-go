// pipeline-filter 模式
package filter

type Request interface{}
type Response interface{}

type IFilter interface {
	Process(data Request) (Response, error)
}
