package sqlconnect

import (
	"context"
	pb "proto/post/gen"
	"utils"
)

//message Comment {
//	string id = 1; // comment id
//	string post_id = 2; // post id
//	string creator_id = 3; // id of the comment creator
//	string content = 4;
//}

//message AddCommentRequest {
//	string post_id = 1;
//	string user_id = 2;
//	string content = 3;
//}

//message AddCommentResponse {
//	string comment_id = 1;
//}

//CREATE TABLE comments (
//	post_id INTEGER NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
//	comment_id INTEGER NOT NULL,
//	user_id INTEGER NOT NULL,
//	content TEXT NOT NULL,
//	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//
//	PRIMARY KEY (post_id, comment_id)
//);

func AddCommentDbHandler(ctx context.Context, req *pb.AddCommentRequest) (string, error) {
	db, err := ConnectDb()
	if err != nil {
		return "", utils.ConnectingToDbError
	}
	defer db.Close()

	query := "INSERT INTO comments (post_id,"
	return "", nil
}

//message GetPostCommentsRequest {
//	string post_id = 1;
//}
//
//message GetPostCommentsResponse {
//	repeated Comment comments = 1;
//}

func GetCommentsDbHandler(ctx context.Context, req *pb.GetPostCommentsRequest) ([]*pb.Comment, error) {
	return nil, nil
}

//message UpdateCommentRequest {
//	string comment_id = 1;
//	string post_id = 2;
//	string updated_content = 3;
//}

func UpdateCommentDbHandler(ctx context.Context, req *pb.UpdateCommentRequest) error {
	return nil
}

//message DeleteCommentRequest {
//	string comment_id = 1;
//	string post_id = 2;
//}

func DeleteCommentDbHandler(ctx context.Context, req *pb.DeleteCommentRequest) error {
	return nil
}
