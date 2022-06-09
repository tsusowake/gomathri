package main

import (
	"fmt"
	"os"
)

const n = 100

func main() {
	var s1, s2 float64 = 0.0, 0.0
	var t float64 = 1.0 - 1.0/(n+1.0)

	// initialized array
	inc, dec := make([]float64, 100), make([]float64, 100)
	for i := 0; i < n; i++ {
		inc[i] = 0.0
		dec[i] = 0.0
	}

	filename := "out.log"
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// print initialized values
	writeToFile(file, []byte(fmt.Sprintf("%f\n", s1)))
	writeToFile(file, []byte(fmt.Sprintf("%f\n", s2)))
	writeToFile(file, []byte(fmt.Sprintf("%f\n", t)))
	writeToFile(file, []byte(fmt.Sprintf("-----\n")))

	for k := 1; k <= 100; k++ {
		s1 += float64(1.0 / (k * (k + 1.0)))
		inc[k-1] = s1 //inc内の順番はkが大きくなる順
	}

	for k := n; k >= 1; k-- {
		s2 += float64(1.0 / (k * (k + 1.0)))
		dec[n-k] = s2 //dec内の順番はkが小さくなる順
	}

	//各kにおけるs値と最終的な誤差を表示
	for k := 0; k < n; k++ {
		// println(fmt.Sprintf("*s値表示"))
		println(fmt.Sprintf("inc[%d] = %e, dec[%d] = %e\n", k, inc[k], k, dec[k]))

		writeToFile(file,
			[]byte(fmt.Sprintf("inc[%d] = %e, dec[%d] = %e\n", k, inc[k], k, dec[k])))

		// printf("*各kにおける誤差の表示\n");
		// printf("dec[%d]_error = %e, dec[%d]_error = %e\n", k, inc[k] - t, k, dec[k] - t);
	}

	writeToFile(file,
		[]byte(fmt.Sprintf("t = %e\n", t)))
	writeToFile(file,
		[]byte(fmt.Sprintf("conclusive_inc_error = %e\n", inc[n-1]-t)))
	writeToFile(file,
		[]byte(fmt.Sprintf("conclusive_inc_error = %e\n", dec[n-1]-t)))

}

func writeToFile(f *os.File, b []byte) error {
	_, err := f.Write(b)
	if err != nil {
		return err
	}
	return nil
}
