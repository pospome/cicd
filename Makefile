#makeの仕様でタスク定義名と同じファイルが存在している場合はタスクが実行されません。
#これを回避するためには、.PHONY: task をタスク定義に付けます。
.PHONY: run
run:
	docker build -t cicd .
	docker run -it --rm --name cicd -p 80:8080 cicd

# -t ... コンテナとホストの出力を接続する
# -i ... コンテナとホストの入力を接続する
# -p host_port:container_port ... ポートのバインド
# --name xxxx ... コンテナに名前を付けれる。
# --rm ... コンテナが exit したら自動的に終了させる

#変数の渡し方
#https://qiita.com/yoskeoka/items/317a3afab370155b3ae8#make実行時に変数の値を渡す

.PHONY: test
test:
	go test ./...
