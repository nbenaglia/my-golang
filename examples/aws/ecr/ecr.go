package ecr

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
)

var svc *ecr.ECR

// RunFlag -
type RunFlag struct {
	Repo        string
	KeepNum     int
	ExcludeTag  arrayFlag
	ListRepo    bool
	CertProfile string
}

type ImageDetailArray []*ecr.ImageDetail

func (s ImageDetailArray) Len() int { return len(s) }
func (s ImageDetailArray) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ImageDetailArray) Less(i, j int) bool {
	t1 := s[i].ImagePushedAt.Unix()
	t2 := s[j].ImagePushedAt.Unix()
	return t1 > t2
}

type arrayFlag []string

func (i *arrayFlag) String() string {
	return fmt.Sprint(*i)
}
func (i *arrayFlag) Set(val string) error {
	*i = append(*i, val)
	return nil
}

var runFlag RunFlag

func init() {
	runFlag = RunFlag{}
	flag.StringVar(&runFlag.Repo, "repo", "", "repository name")
	flag.IntVar(&runFlag.KeepNum, "keep", 5, "keep number images")
	flag.Var(&runFlag.ExcludeTag, "e", "exclude delete tag")
	flag.BoolVar(&runFlag.ListRepo, "list-repo", false, "list all repository")
	flag.StringVar(&runFlag.CertProfile, "profile", "", "aws config profile")
	flag.Parse()
}

func main() {
	var cred *credentials.Credentials
	_ = cred
	region := os.Getenv("AWS_REGION")
	if len(region) == 0 {
		region = "eu-west-1"
	}
	if len(runFlag.CertProfile) == 0 {
		accessKey := os.Getenv("AWS_ACCESS_KEY")
		secretKey := os.Getenv("AWS_SECRET_KEY")
		if len(accessKey) == 0 || len(secretKey) == 0 {
			log.Fatal("AWS_ACCESS_KEY or AWS_SECRET_KEY env not found")
		}
		cred = credentials.NewStaticCredentials(accessKey, secretKey, "")
	} else {
		cred = credentials.NewSharedCredentials("", runFlag.CertProfile)
	}
	conf := &aws.Config{
		Region:      aws.String(region),
		Credentials: cred,
	}
	sess, _ := session.NewSession(conf)
	svc = ecr.New(sess)

	if runFlag.ListRepo == true {
		repos := showRepo()
		for k, v := range repos.Repositories {
			fmt.Printf("%d) %s\n", k, *v.RepositoryName)
		}
		return
	}

	if len(runFlag.Repo) == 0 {
		log.Fatal("repository name is empty")
	}

	repos := showRepo(runFlag.Repo)
	if len(repos.Repositories) == 0 {
		log.Fatal("repository not found")
	}

	images := showImages(runFlag.Repo)
	if len(images.ImageDetails) == 0 {
		fmt.Println("no image found")
		return
	}

	if len(images.ImageDetails) <= runFlag.KeepNum {
		fmt.Println("image count <= keep")
		return
	}

	detailArray := ImageDetailArray{}
	keepNum := runFlag.KeepNum
	for _, v := range images.ImageDetails {
		excludeFlag := false
		if len(runFlag.ExcludeTag) > 0 {
			for _, v2 := range v.ImageTags {
				f := false
				for _, e := range runFlag.ExcludeTag {
					if *v2 == e {
						f = true
						excludeFlag = true
						keepNum--
						break
					}
				}
				if f == true {
					break
				}
			}
		}
		if !excludeFlag {
			detailArray = append(detailArray, v)
		}
	}
	if keepNum < 0 {
		keepNum = 0
	}
	sort.Stable(detailArray)
	detailArray = detailArray[keepNum:]
	if len(detailArray) == 0 {
		fmt.Println("no delete image")
		return
	}

	// show delete image info
	for k, v := range detailArray {
		tags := make([]string, len(v.ImageTags))
		for idx, vv := range v.ImageTags {
			tags[idx] = *vv
		}
		fmt.Printf("%d) %s (%s)\n", k, *v.ImageDigest, strings.Join(tags, ","))
	}
	txt := readInput("Delete (Y/n)")
	if len(txt) == 0 || regexp.MustCompile(`(?i)^y$`).Match([]byte(txt)) {
		// do delete
		fmt.Println("do delete")
		tmp := make([]string, len(detailArray))
		for k, v := range detailArray {
			tmp[k] = *v.ImageDigest
		}
		deleteImages(runFlag.Repo, tmp)
	} else {
		fmt.Println("cancel!")
		return
	}

	fmt.Println("finish")
}

func readInput(tip string) string {
	fmt.Printf("%s : ", tip)
	reader := bufio.NewReader(os.Stdin)
	txt, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("input text ====> ", txt)
	txt = strings.Replace(txt, "\n", "", -1)
	txt = strings.Trim(txt, " ")
	return txt
}

func showRepo(n ...interface{}) *ecr.DescribeRepositoriesOutput {
	in := &ecr.DescribeRepositoriesInput{}

	if len(n) > 0 {
		name, ok := n[0].(string)
		if ok {
			in.RepositoryNames = []*string{aws.String(name)}
		}
	}

	out, err := svc.DescribeRepositories(in)
	if err != nil {
		log.Fatal(err)
	}
	return out
}

func showImages(repo string) *ecr.DescribeImagesOutput {
	if len(repo) == 0 {
		fmt.Println("repository name is empty")
		return nil
	}
	in := &ecr.DescribeImagesInput{
		RepositoryName: aws.String(repo),
	}
	out, err := svc.DescribeImages(in)
	if err != nil {
		log.Fatal(err)
	}
	return out
}

func deleteImages(repo string, digest []string) {
	if len(repo) == 0 {
		return
	}
	if len(digest) == 0 {
		return
	}

	images := make([]*ecr.ImageIdentifier, len(digest))
	for k, v := range digest {
		t := &ecr.ImageIdentifier{}
		t.ImageDigest = aws.String(v)
		images[k] = t
	}

	in := &ecr.BatchDeleteImageInput{}
	in.SetRepositoryName(repo)
	in.SetImageIds(images)

	_, err := svc.BatchDeleteImage(in)
	if err != nil {
		log.Fatal(err)
	}
}
