package main

import (
	"fmt"
	"mu_previous_papers_be/model"
	"mu_previous_papers_be/server"
	"mu_previous_papers_be/store"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	dns := os.Getenv("DATABASE_URL")
	db, err := model.NewDB(dns)
	if err != nil {
		fmt.Println("error connecting to pgsql server: ", err)
		return
	}
	str := store.NewStore(db)
	serverMain := server.NewServer(str)
	serverMain.Run()

	// var bucketName = "previous-papers-mu"
	// var accountId = "27d217361f3af8a6b0d8330ce86a4b8c"
	// var accessKeyId = "6c39077fc587570d9b7bcbb8c201afe3"
	// var accessKeySecret = "d73520d6e102d50d558215cf825c64696f8c0b7a285739685c5a5f76e8adcbd1"

	// r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
	// 	return aws.Endpoint{
	// 		URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountId),
	// 	}, nil
	// })

	// cfg, err := config.LoadDefaultConfig(context.TODO(),
	// 	config.WithEndpointResolverWithOptions(r2Resolver),
	// 	config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
	// 	config.WithRegion("auto"),
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// client := s3.NewFromConfig(cfg)

	// listObjectsOutput, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
	// 	Bucket: &bucketName,
	// })
	// if err != nil {
	// 	log.Fatal("ress in llll: ", err)
	// }

	// for _, object := range listObjectsOutput.Contents {
	// 	obj, _ := json.MarshalIndent(object, "", "\t")
	// 	fmt.Println(string(obj))
	// }

	// listBucketsOutput, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, object := range listBucketsOutput.Buckets {
	// 	obj, _ := json.MarshalIndent(object, "", "\t")
	// 	fmt.Println(string(obj))
	// }

	// localFilePath := "local-file.pdf"
	// downloadFile, err := os.Create(localFilePath)
	// if err != nil {
	// 	fmt.Println("error here ttt: ", err)
	// }
	// defer downloadFile.Close()

	// downloader := manager.NewDownloader(client)
	// numBytes, err := downloader.Download(context.TODO(), downloadFile, &s3.GetObjectInput{
	// 	Bucket: aws.String(bucketName),
	// 	Key:    aws.String("ES303_Major_sem6_2019.pdf"),
	// })
	// if err != nil {
	// 	fmt.Println("error here jjjj: ", err)
	// }
	// fmt.Println("numbytes: ", numBytes)

}
