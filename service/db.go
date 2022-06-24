package service

import (
	"github.com/1920853199/passwd/codec"
	"github.com/dgraph-io/badger/v3"
	"github.com/dgraph-io/badger/v3/options"
)

type DB struct {
	db *badger.DB
}

var (
	pb = codec.MsgpackCodec{}
)

func NewDB(path string) (*DB, error) {
	db, err := badger.Open(badger.DefaultOptions(path).
		WithCompression(options.ZSTD).
		WithSyncWrites(false).
		WithBlockCacheSize(100 * (1 << 20)).
		WithIndexCacheSize(100 * (1 << 20)).
		WithZSTDCompressionLevel(3))

	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func (db *DB) Set(key, value []byte) error {
	fn := func(tx *badger.Txn) error {
		return tx.Set(key, (value))
	}
	return db.db.Update(fn)
}

func (db *DB) Get(key []byte) ([]byte, error) {
	var val []byte
	fn := func(tx *badger.Txn) error {
		item, err := tx.Get(key)
		if err != nil {
			return err
		}
		val, _ = item.ValueCopy(nil)
		return nil
	}

	err := db.db.View(fn)
	return val, err
}

func (db *DB) SetWithCodec(key []byte, value interface{}) error {
	data, err := pb.Encode(value)
	if err != nil {
		return err
	}
	return db.Set(key, data)
}

func (db *DB) GetWithCodec(key []byte, value interface{}) error {
	data, err := db.Get(key)
	if err != nil {
		return err
	}
	return pb.Decode(data, value)
}

// 删除KEY
func (db *DB) Del(key []byte) error {
	fn := func(tx *badger.Txn) error {
		return tx.Delete(key)
	}
	return db.db.Update(fn)
}

// 清理全部数据
func (db *DB) Clear() error {
	fn := func(tx *badger.Txn) error {
		opt := badger.DefaultIteratorOptions
		opt.Reverse = true
		opt.PrefetchValues = false
		it := tx.NewIterator(opt)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			tx.Delete(it.Item().KeyCopy(nil))
		}
		return nil
	}

	return db.db.Update(fn)
}

// 查看全部数据
func (db *DB) All() (interface{}, error) {
	txn := db.db.NewTransaction(false)
	defer txn.Discard()
	iter := badger.DefaultIteratorOptions
	it := txn.NewIterator(iter)
	defer it.Close()
	var res = make([]Args, 0)
	for it.Rewind(); it.Valid(); it.Next() {
		item := it.Item()

		if string(item.Key()) == TOKENKEY {
			continue
		}

		value, _ := item.ValueCopy(nil)

		var tmp Item
		err := pb.Decode(value, &tmp)
		if err != nil {
			return nil, err
		}

		res = append(res, Args{
			Key:   string(item.Key()),
			Value: tmp,
		})
	}

	return res, nil
}
