package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/user"
)

var photourls = []string{"http://i.ytimg.com/vi/8MtQRv6aKx4/maxresdefault.jpg",
	"https://pmchollywoodlife.files.wordpress.com/2014/07/taylor-swift-cat-lady-3.jpg?w=600",
	"http://afrossip.com/wp-content/uploads/2015/03/95fb4cfc9b9b4741ba92fdab96ccf708.ashx_.jpeg",
	"https://pbs.twimg.com/media/Bv1Wb-kIAAAb7RW.jpg",
	"http://www.eonline.com/eol_images/Entire_Site/2014915/rs_560x415-141015174855-1024.Taylor-Swift-DIET-COKE-CAT-COMMERCIAL.MS.101514_copy.JPG",
	"http://media3.popsugar-assets.com/files/2015/06/24/759/n/1922398/139fba6e_edit_img_cover_file_14690959_1433442600_10748096_31083IP9A3Q.xxxlarge/i/Pictures-Taylor-Swift-Cats.jpg",
	"https://sellmeaboutit.files.wordpress.com/2014/11/taylor-swift-cat-5.jpg",
	"https://s-media-cache-ak0.pinimg.com/236x/52/de/c9/52dec9fe356ae4168ecbc6a887da4a61.jpg",
	"https://pbs.twimg.com/media/Bv1Wb-kIAAAb7RW.jpg",
	"http://images6.fanpop.com/image/photos/36400000/rids-and-me-image-rids-and-me-36451428-225-225.jpg"}

func main() {
	usr, err := user.Current()
	if err != nil {
		fmt.Printf("%v", err)
	}
	err = os.Chdir(fmt.Sprintf("%v/Desktop", usr.HomeDir))
	if err != nil {
		fmt.Printf("%v", err)
	}
	client := http.Client{}
	iterations := 10
	photos := len(photourls)
	for i := 0; i < photos*iterations; i++ {
		// urlN := photourls[i%photos]
		// fileURL, err := url.Parse(urlN)
		// path := fileURL.Path
		//segments := strings.Split(path, "/")
		filename := fmt.Sprintf("%vCatswift", i)

		file, err := os.Create(filename)
		file.Chdir()
		defer file.Close()
		img, err := client.Get(photourls[i%photos])
		if err != nil {
			fmt.Println("error")
			continue
		}
		defer img.Body.Close()

		_, err = io.Copy(file, img.Body)
		if err != nil {
			fmt.Printf("error%v\n", err)
		}
	}
}
