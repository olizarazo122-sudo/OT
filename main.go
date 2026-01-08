package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	//"io"
	"bufio"
	"log"
	"os/exec"
	"runtime"
	"strings"

	"github.com/logrusorgru/aurora"
)

var ops string
var devices = []string{}
var fs1 sync.WaitGroup
var fs2 sync.WaitGroup
var fs3 sync.WaitGroup
var adbss string

func main() {
	fmt.Println("###############################################")
	fmt.Println(aurora.Green("# Bienvenido a la herramienta de optimizacion #"))
	fmt.Println("")
	fmt.Println(aurora.Yellow("Por favor tener en cuenta lo siguiente"))
	fmt.Println(aurora.Yellow("esta herramienta fue creada especificamente para los dispositivos 'HTG 1078-GOB con SoC RK3066' "))
	fmt.Println(aurora.Yellow("esta herramienta fue creada para Android 4.4"))
	fmt.Println(aurora.Yellow("si lo usa con otro modelo puede causar efectos no conocidos"))
	fmt.Println(aurora.Red("EL USO DE ESTA HERRAMIENTA SERA TOTALMENTE BAJO SU RESPONSABILIDAD"))
	fmt.Println(aurora.Red("Algunos APK pueden no ser compatibles, por favor verificar y reportar"))
	fmt.Println("")
	fmt.Println(aurora.Green("acontinuación por favor seleccione una opción"))
	fmt.Println()

	fmt.Println(aurora.Red("para Hacer uso de esta herramienta hay varias condiciones"))
	fmt.Println("1. tener el modo desarrollador activado junto a la depuración usb")
	fmt.Println("2. la tablet en su estado de fabrica reseteada")

	fmt.Println()
	fmt.Println("El modo desarrollador de puede acceder llendo a Ajustes-Información de la tablet,y dando click 7 veces sobre el numero de compilación")
	fmt.Println("para resetearlo puedes ir a ajustes-copia de seguridad-restablecer datos de fabrica y dando click en restablecer tablet con el borrado de la SD activado")

	fmt.Println("por favor precione Y para seguir y N para cancelar si las condiciones no se han cumplido")
	var state string
	fmt.Scanln(&state)
	fmt.Println(state)

	if runtime.GOOS == "windows" {
		adbss = "./adb.exe"
	} else {
		adbss = "./adb"
	}

	clear()

	fmt.Println()
	fmt.Println(aurora.Red("el proceso se puede ejecutar hasta con 3 tablets simultaneamente"))
	fmt.Println(aurora.Red("Por favor seleccione el/los dispositivos"))

	selectDevices()

	if state == "Y" || state == "y" {
		fmt.Println(aurora.Green("Iniciando"))

		cantdevices := len(devices)
		for i := 0; i < 3; i++ {
			if i == 0 {
				fs1.Add(cantdevices)
				for e := 0; e < cantdevices; e++ {
					go optimdevice(devices[e])
				}
				fs1.Wait()
			} else if i == 1 {
				fs2.Add(cantdevices)
				for e := 0; e < cantdevices; e++ {
					go installapp(devices[e])
				}
				fs2.Wait()
			} else if i == 2 {
				fs3.Add(cantdevices)
				for e := 0; e < cantdevices; e++ {
					go uninstallapp(devices[e])
				}
				fs3.Wait()
			}
		}

		fmt.Println(aurora.Green("El proceso termino"))
		fmt.Println(aurora.Yellow("Por favor reinicie la tablet"))
		time.Sleep(6 * time.Second)

	} else {
		fmt.Println(aurora.Red("Cerrando"))
		return
	}
}

func selectDevices() {
	cmd := exec.Command("./adb", "devices", "-l")

	if runtime.GOOS == "linux" {
		cmd.Dir = "adb/linux"
	} else if runtime.GOOS == "darwin" {
		cmd.Dir = "adb/mac"
	} else if runtime.GOOS == "windows" {
		cmd.Dir = "adb/windows"
	}

	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile("adb.txt", out, 0644)

	file, err := os.Open("adb.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		linea := scanner.Text()
		device := strings.SplitN(linea, " ", 7)

		if len(device) == 7 {
			serial := device[0]

			modelo := device[5]
			dispositivo := device[6]

			fmt.Print(modelo)
			fmt.Print("  ")
			fmt.Print(dispositivo)
			fmt.Println("")

			var dv string
			fmt.Println("Si este dispositivo se optimizara por favor marcar y/n")
			fmt.Scan(&dv)

			if dv == "Y" || dv == "y" {
				devices = append(devices, serial)
			}

		}
	}

	os.Remove("adb.txt")

}

func optimdevice(device string) {

	defer fs1.Done()
	clear()

	dt := time.Now()
	time := dt.Format("20060102/150405")

	cmd := exec.Command("adb")
	if runtime.GOOS == "linux" {
		cmd.Dir = "adb/linux"
	} else if runtime.GOOS == "darwin" {
		cmd.Dir = "adb/mac"
	} else if runtime.GOOS == "windows" {
		cmd.Dir = "adb/windows"
	}

	cmd = exec.Command(adbss, "-s", device, "shell", "settings", "put", "global", "window_animation_scale", "0")
	cmd.Run()
	cmd = exec.Command(adbss, "-s", device, "shell", "settings", "put", "global", "window_animation_scale", "0")
	cmd.Run()
	cmd = exec.Command(adbss, "-s", device, "shell", "settings", "put", "global", "transition_animation_scale", "0")
	cmd.Run()
	cmd = exec.Command(adbss, "-s", device, "shell", "settings", "put", "global", "animator_duration_scale", "0")
	cmd.Run()
	cmd = exec.Command(adbss, "-s", device, "shell", "settings", "put", "global", "background_process_limit", "3")
	cmd.Run()
	cmd = exec.Command(adbss, "-s", device, "shell", "settings", "put", "global", "force_hw_ui", "1")
	cmd.Run()
	cmd = exec.Command(adbss, "-s", device, "shell", "settings", "put", "global", "mobile_data_always_on", "0")
	cmd.Run()
	cmd = exec.Command(adbss, "-s", device, "shell", "date", time)
	cmd.Run()

}

func clear() {

	cmd := exec.Command("adb") //provisional

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

}

func installapp(device string) {
	defer fs2.Done()
	clear()

	file, err := os.Open("install.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		linea := scanner.Text()

		operacion := strings.SplitN(linea, " ", 3)

		if len(operacion) == 3 {
			app := operacion[0]
			install := operacion[1]

			fmt.Print(app)
			fmt.Print(" " + install)
			fmt.Println()

			if install == "true" {
				path := "../../apks/" + app + ".apk"
				path = strings.TrimSpace(path)

				cmd := exec.Command(adbss, "-s", device, "install", path)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				if runtime.GOOS == "linux" {
					cmd.Dir = "adb/linux"
				} else if runtime.GOOS == "darwin" {
					cmd.Dir = "adb/mac"
				} else if runtime.GOOS == "windows" {
					cmd.Dir = "adb/windows"
				}

				cmd.Run()
			}

		}

	}

}

func uninstallapp(device string) {

	defer fs3.Done()
	file, err := os.Open("uninstall.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		linea := scanner.Text()

		operacion := strings.SplitN(linea, " ", 3)

		if len(operacion) == 3 {
			app := operacion[0]
			uninstall := operacion[1]
			action := operacion[2]

			fmt.Print(app)
			fmt.Print(" " + uninstall)
			fmt.Println()

			cmd := exec.Command("adb") //provisional

			if action == "U" {
				cmd = exec.Command(adbss, "-s", device, "uninstall", app)
			} else if action == "D" {
				cmd = exec.Command(adbss, "-s", device, "shell", "pm", "disable-user", "--user", "0", app)
			}

			if runtime.GOOS == "linux" {
				cmd.Dir = "adb/linux"
			} else if runtime.GOOS == "darwin" {
				cmd.Dir = "adb/mac"
			} else if runtime.GOOS == "windows" {
				cmd.Dir = "adb/windows"
			}

			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()

		}

	}

}
