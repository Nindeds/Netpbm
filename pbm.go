package Netpbm

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PBM struct {
	data          [][]bool
	width, height int
	magicNumber   string
}

/* main() {
	pbmCall, err := ReadPBM("Image.pbm")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier PBM:", err)
		return
	}
	display(pbmCall.data)
}

*/

func ReadPBM(filename string) (*PBM, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var width, height int
	//Create a scanner that read line by line
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	magicNumber := scanner.Text() // Take last token of scanner scan
	if magicNumber != "P1" && magicNumber != "P4" {
		return nil, errors.New("unsupported file type")
	}
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "#") {
			_, err := fmt.Sscanf(line, "%d %d", &width, &height)
			if err == nil {
				break
			} else {
				fmt.Println("Invalid width or height:", err)
			}
		}
	}
	var data [][]bool
	for scanner.Scan() {
		line := scanner.Text()
		if magicNumber == "P1" {
			row := make([]bool, width)
			for i, char := range strings.Fields(line) {
				pixel, _ := strconv.Atoi(char) //Convert (char) to int
				if pixel == 1 {
					row[i] = true
				}
			}
			data = append(data, row)
		} else if magicNumber == "P4" {

		}
	}

	return &PBM{
		data:        data,
		width:       width,
		height:      height,
		magicNumber: magicNumber,
	}, nil
}

func (pbm *PBM) Size() (int, int) {
	return pbm.width, pbm.height
}

func (pbm *PBM) At(x, y int) bool {
	if x >= 0 && x < pbm.height && y >= 0 && y < pbm.width {
		return pbm.data[x][y]
	}
	return false
}
func (pbm *PBM) Set(x, y int, value bool) {
	pbm.data[x][y] = value
}

// saves the PBM image to a file and returns an error if there was a problem (only p1)
func (pbm *PBM) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier:", err)
		return err
	}
	defer file.Close()

	fmt.Fprintf(file, "%s\n%d %d\n", pbm.magicNumber, pbm.width, pbm.height)

	for _, row := range pbm.data {
		for _, pixel := range row {
			if pixel {
				fmt.Print(file, "1 ")
			} else {
				fmt.Print(file, "0 ")
			}
		}
		fmt.Println(file)
	}
	fmt.Println("Données écrites avec succès dans le fichier PBM.")
	return nil
}

// Inverts  the colors of the PBM image.
func (pbm *PBM) Invert() {
	for i := 0; i < len(pbm.data); i++ {
		for j := 0; j < len(pbm.data[i]); j++ {
			pbm.data[i][j] = !pbm.data[i][j]
		}
	}
}

// It flip the PBM image horizontally
func (pbm *PBM) Flip() {
	NumRows := pbm.width
	Numcolums := pbm.height
	for i := 0; i < NumRows; i++ {
		for j := 0; j < Numcolums/2; j++ {
			pbm.data[i][j], pbm.data[i][Numcolums-j-1] = pbm.data[i][Numcolums-j-1], pbm.data[i][j]
		}
	}
}

// flops the PBM image vertically
func (pbm *PBM) Flop() {
	numRows := len(pbm.data)
	if numRows == 0 {
		return
	}
	for i := 0; i < numRows/2; i++ {
		pbm.data[i], pbm.data[numRows-i-1] = pbm.data[numRows-i-1], pbm.data[i]
	}
}

// Set new magic number
func (pbm *PBM) SetMagicNumber(magicNumber string) {
	pbm.magicNumber = magicNumber
	if magicNumber != "P1" && magicNumber != "P4" {
		fmt.Println("unsupported format")
	} else {
		fmt.Println(magicNumber)
	}
}
