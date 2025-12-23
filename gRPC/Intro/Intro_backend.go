package intro

import (
	pb "Intro/Intro_pb"
	"context"
	"errors"
	"log"
)


type Info struct {
	Result map[string]*pb.Codename
}


func (I *Info) AddCodeName(ctx context.Context, Info *pb.VerInfo) (*pb.Nil, error) {
	if _,ok:=I.Result[Info.Ab.Version]; !ok{
		I.Result[Info.Ab.Version]= &pb.Codename{
			 Codename: Info.Cn.Codename,
			 Package: Info.Cn.Package,
		}
		log.Printf("New Entry for %v", Info.Ab.Version)
		log.Println()
	}else{
		log.Println("Already Exists!")
		return &pb.Nil{}, errors.New("Version already exists")
	}
	return &pb.Nil{}, nil
}
