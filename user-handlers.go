package main

import (
	"encoding/json"
	"io"
	//"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//UserCreate ...
func UserCreate(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Access-Control-Allow-Origin", "*")

	msg := new(AddUserMsg)
	FillStruct(r, msg)

	insertResp := InsertUser(*msg)
	setUpUserDirectory(insertResp.ID)

	// w.WriteHeader(http.StatusCreated)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	ifErr(json.NewEncoder(w).Encode(insertResp))
}

//UserDelete ...
func UserDelete(w http.ResponseWriter, r *http.Request) {
	msg := new(DelUserMsg)
	FillStruct(r, msg)
	RemoveUser(*msg)
	ifErr(json.NewEncoder(w).Encode("Completed"))
}

//Login ...
func Login(w http.ResponseWriter, r *http.Request) {
	msg := new(LoginMsg)
	FillStruct(r, msg)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	ifErr(json.NewEncoder(w).Encode(RetrieveUser(*msg)))
}

//ProfPic ...
func ProfPic(w http.ResponseWriter, r *http.Request) {
	UID := mux.Vars(r)["UID"]
	// TODO: check if UID is a valid UID
	f, err := os.Open(PrjDir + UID + "/" + UID)
	ifErr(err)
	io.Copy(w, f)
	defer f.Close()
}

//GetFriends ...
func GetFriends(w http.ResponseWriter, r *http.Request) {
	UID := mux.Vars(r)["UID"]
	w.Header().Set("Access-Control-Allow-Origin", "*")
	ifErr(json.NewEncoder(w).
		Encode(
			GetFriendsResponse{GetFriendReqs(UID), GetFriendsMgo(UID)}))
}

//SearchUser ...
func SearchUser(w http.ResponseWriter, r *http.Request) {
	nameLike := mux.Vars(r)["NAME_LIKE"]
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fechedProfiles := GetProfilesMgo(nameLike)
	ifErr(json.NewEncoder(w).Encode(fechedProfiles))
}

// elseif(!empty($_POST['ACPT_REQ'])){
// 	$uid = new MongoDB\BSON\ObjectId($_POST['UID']);
// 	$friendId = new MongoDB\BSON\ObjectId($_POST['FRND_ID']);
// 	$collection->updateOne(
// 		array( "_id" => $friendId),
// 		array( '$pull'=> array("friendReqs"=>$uid))
// 	);
// 	$result = $collection->updateOne(
// 		array( "_id" => $uid),
// 		array( '$pull'=> array("friendReqs"=>$friendId))
// 	);
// 	if($result->getModifiedCount() != 1){
// 		echo json_encode(array("err"=>true, "msg" => "Failed to remove from reqs list."));
// 		exit(0);
// 	}
// 	else{
// 		$result = $collection->updateOne(
// 			array( "_id" => $friendId),
// 			array( '$addToSet'=> array("friends"=>$uid))
// 		);
// 		if($result->getModifiedCount() != 1){
// 			echo json_encode(array("err"=>true, "msg" => "Failed to add id to friends list."));
// 			exit(0);
// 		}
// 		$result = $collection->updateOne(
// 			array( "_id" => $uid),
// 			array( '$addToSet'=> array("friends"=>$friendId))
// 		);
// 		if($result->getModifiedCount() != 1){
// 			echo json_encode(array("err"=>true, "msg" => "Failed to add id to friends list."));
// 			exit(0);
// 		}
// 		echo json_encode(array("err"=>false));
// 		exit(0);
// 	}
// }

// elseif(!empty($_POST['DEC_REQ'])){
// 	$uid = new MongoDB\BSON\ObjectId($_POST['UID']);
// 	$friendId = new MongoDB\BSON\ObjectId($_POST['FRND_ID']);
// 	$result = $collection->updateOne(
// 		array( "_id" => $uid),
// 		array( '$pull'=> array("friendReqs"=>$friendId))
// 	);
// 	if($result->getModifiedCount() != 1){
// 		echo json_encode(array("err"=>true, "msg" => "Failed to delete album."));
// 		exit(0);
// 	}
// 	else{
// 		echo json_encode(array("err"=>false));
// 		exit(0);
// 	}
// }
