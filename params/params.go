/**
 * @author: chenyang/904852749@qq.com
 * @doc:
 **/

package params

import (
	"context"
	"log"
	"sync"
)

const CustomBizKey = "biz"

func InitContextParams(ctx context.Context) context.Context {
	return context.WithValue(ctx, CustomBizKey, &sync.Map{})
}

func AppendContextParams(ctx context.Context, values map[string]interface{}) {
	params, ok := ctx.Value(CustomBizKey).(*sync.Map)
	if !ok {
		log.Printf("[params.AppendContextParams] Fails to get params from context")
		return
	}

	for k, v := range values {
		params.Store(k, v)
	}
}

func GetTypedContextParams[T any](ctx context.Context, mapKey string) (T, bool) {
	var empty T
	value, ok := getContextParams(ctx, mapKey)
	if !ok {
		return empty, false
	}

	valueT, ok := value.(T)
	if !ok {
		return empty, false
	}

	return valueT, true
}

func MustGetContextParams[T any](ctx context.Context, mapKey string) T {
	var empty T
	value, ok := getContextParams(ctx, mapKey)
	if !ok {
		log.Printf("[params.MustGetContextParams] cannot get key: %v", mapKey)
		return empty
	}
	valueT, ok := value.(T)
	if !ok {
		log.Printf("[params.MustGetContextParams] value not string, key: %v", mapKey)
		return empty
	}
	return valueT
}

func getContextParams(ctx context.Context, mapKey string) (interface{}, bool) {
	params, ok := ctx.Value(CustomBizKey).(*sync.Map)
	if !ok {
		log.Printf("[params.GetContextParams] Failed to get params from context")
		return nil, false
	}

	return params.Load(mapKey)
}
