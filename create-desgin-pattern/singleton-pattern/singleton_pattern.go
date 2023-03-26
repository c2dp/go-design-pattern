package singletonpattern

import "sync"

type sigleton struct{}

var instance *sigleton

var mtx sync.Mutex

// 懒汉式单例模式
func GetInstanceLazy() *sigleton {
	if instance == nil {
		instance = new(sigleton)
	}
	return instance
}

// 懒汉式单例模式(协程安全)
func GetInstanceLazySafe() *sigleton {
	mtx.Lock()
	defer mtx.Unlock()
	if instance == nil {
		instance = new(sigleton)
	}
	return instance
}

// 饿汉式单例模式（协程安全）
func init() {
	if instance == nil {
		instance = new(sigleton)
	}
}

// 双重校验单例模式
func GetInstanceDoubleCheck() *sigleton {
	if instance == nil {
		mtx.Lock()
		if instance == nil {
			instance = new(sigleton)
		}
		mtx.Unlock()
	}
	return instance
}

// sync.Once 单例模式 推荐

var once sync.Once

func GetInstanceBySyncOnce() *sigleton {
	once.Do(func() {
		instance = new(sigleton)

	})
	return instance
}
