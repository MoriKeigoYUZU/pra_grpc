package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"pra_grpc/chat"
)

func main() {
	lis, err := net.Listen("tcp", ":9000") // TCPリスナーをポート9000で作成
	if err != nil {                        // エラーチェック
		log.Fatalf("failed to listen: %v", err) // リスナー作成失敗時にログ出力してプログラムを終了
	}

	s := chat.Server{}                             // チャットサーバーのインスタンスを作成
	grpcServer := grpc.NewServer()                 // 新しいgRPCサーバーを作成
	chat.RegisterChatServiceServer(grpcServer, &s) // チャットサービスをgRPCサーバーに登録
	if err := grpcServer.Serve(lis); err != nil {  // gRPCサーバーを開始してリクエストを受け付ける
		log.Fatalf("failed to serve: %s", err) // サーバー起動失敗時にログ出力してプログラムを終了
	}
}
