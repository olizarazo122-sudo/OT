package main

import (
	"fmt"
	"os"
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

	if state == "Y" || state == "y" {
		fmt.Println(aurora.Green("Iniciando"))
		installapp()
		uninstallapp()
		optimdevice()
		clear()

		fmt.Println(aurora.Green("El proceso termino"))
		fmt.Println(aurora.Yellow("Por favor reinicie la tablet"))
		time.Sleep(6 * time.Second)

	} else {
		fmt.Println(aurora.Red("Cerrando"))
		return
	}
}

func optimdevice() {
	clear()

	if runtime.GOOS == "linux" {
		cmd := exec.Command("./adb", "shell", "settings", "put", "global", "window_animation_scale", "0")
		cmd.Dir = "src/adb-linux"
		cmd.Run()
		cmd_one := exec.Command("./adb", "shell", "settings", "put", "global", "transition_animation_scale", "0")
		cmd_one.Dir = "src/adb-linux"
		cmd_one.Run()
		cmd_two := exec.Command("./adb", "shell", "settings", "put", "global", "animator_duration_scale", "0")
		cmd_two.Dir = "src/adb-linux"
		cmd_two.Run()
		cmd_three := exec.Command("./adb", "shell", "settings", "put", "global", "background_process_limit", "3")
		cmd_three.Dir = ("src/adb-linux")
		cmd_three.Run()
		cmd_for := exec.Command("./adb", "shell", "settings", "put", "global", "force_hw_ui", "1")
		cmd_for.Dir = ("src/adb-linux")
		cmd_for.Run()
		cmd_five := exec.Command("./adb", "shell", "settings", "put", "global", "mobile_data_always_on", "0")
		cmd_five.Dir = ("src/adb-linux")
		cmd_five.Run()
		cmd_t := exec.Command("./adb", "shell", "date", "%m%d%H%M%Y.%S")
		cmd_t.Dir = "src/adb-linux"
		cmd_t.Run()

	} else if runtime.GOOS == "windows" {

		cmd := exec.Command("./adb.exe", "shell", "settings", "put", "global", "window_animation_scale", "0")
		cmd.Dir = "src/adb-windows"
		cmd.Run()
		cmd_one := exec.Command("./adb.exe", "shell", "settings", "put", "global", "transition_animation_scale", "0")
		cmd_one.Dir = "src/adb-windows"
		cmd_one.Run()
		cmd_two := exec.Command("./adb.exe", "shell", "settings", "put", "global", "animator_duration_scale", "0")
		cmd_two.Dir = "src/adb-windows"
		cmd_two.Run()
		cmd_three := exec.Command("./adb.exe", "shell", "settings", "put", "global", "background_process_limit", "3")
		cmd_three.Dir = ("src/adb-windows")
		cmd_three.Run()
		cmd_for := exec.Command("./adb.exe", "shell", "settings", "put", "global", "force_hw_ui", "1")
		cmd_for.Dir = ("src/adb-windows")
		cmd_for.Run()
		cmd_five := exec.Command("./adb.exe", "shell", "settings", "put", "global", "mobile_data_always_on", "0")
		cmd_five.Dir = ("src/adb-windows")
		cmd_five.Run()
		cmd_t := exec.Command("adb,exe", "shell", "date", "%m%d%H%M%Y.%S")
		cmd_t.Dir = "src/adb-windows"
		cmd_t.Run()
	} else if runtime.GOOS == "darwin" {
		cmd := exec.Command("./adb", "shell", "settings", "put", "global", "window_animation_scale", "0")
		cmd.Dir = "src/adb-mac"
		cmd.Run()
		cmd_one := exec.Command("./adb", "shell", "settings", "put", "global", "transition_animation_scale", "0")
		cmd_one.Dir = "src/adb-mac"
		cmd_one.Run()
		cmd_two := exec.Command("./adb", "shell", "settings", "put", "global", "animator_duration_scale", "0")
		cmd_two.Dir = "src/adb-mac"
		cmd_two.Run()
		cmd_three := exec.Command("./adb", "shell", "settings", "put", "global", "background_process_limit", "3")
		cmd_three.Dir = ("src/adb-mac")
		cmd_three.Run()
		cmd_for := exec.Command("./adb", "shell", "settings", "put", "global", "force_hw_ui", "1")
		cmd_for.Dir = ("src/adb-mac")
		cmd_for.Run()
		cmd_five := exec.Command("./adb", "shell", "settings", "put", "global", "mobile_data_always_on", "0")
		cmd_five.Dir = ("src/adb-mac")
		cmd_five.Run()
		cmd_t := exec.Command("./adb", "shell", "date", "%m%d%H%M%Y.%S")
		cmd_t.Dir = "src/adb-mac"
		cmd_t.Run()

	}

}

func clear() {
	fmt.Println("limpiando consola")
	fmt.Print("\033[2J")

}

func installapp() {
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

			if runtime.GOOS == "linux" {
				if install == "true" {

					path := "../../apks/" + app + ".apk"
					path = strings.TrimSpace(path)

					cmd := exec.Command("./adb", "install", path)
					cmd.Dir = "src/adb-linux"
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					cmd.Run()

				}
			} else if runtime.GOOS == "windows" {
				if install == "true" {

					path := "../../apks/" + app + ".apk"
					path = strings.TrimSpace(path)

					cmd := exec.Command("./adb.exe", "install", path)
					cmd.Dir = "src/adb-windows"
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					cmd.Run()

				}
			} else if runtime.GOOS == "darwin" {

				if install == "true" {

					path := "../../apks/" + app + ".apk"
					path = strings.TrimSpace(path)

					cmd := exec.Command("./adb", "install", path)
					cmd.Dir = "src/adb-mac"
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					cmd.Run()

				}
			}

		}

	}

}

func uninstallapp() {

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

			if uninstall == "true" {
				if runtime.GOOS == "linux" {

					if action == "U" {
						cmd := exec.Command("./adb", "uninstall", app)
						cmd.Dir = "src/adb-linux"
						cmd.Stdout = os.Stdout
						cmd.Stderr = os.Stderr
						cmd.Run()
					} else {
						cmd := exec.Command("./adb", "shell", "pm", "disable-user", "--user", "0", app)
						cmd.Dir = "src/adb-linux"
						cmd.Stdout = os.Stdout
						cmd.Stderr = os.Stderr
						cmd.Run()
					}

				} else if runtime.GOOS == "windows" {

					if action == "U" {
						cmd := exec.Command("./adb.exe", "uninstall", app)
						cmd.Dir = "src/windows"
						cmd.Stdout = os.Stdout
						cmd.Stderr = os.Stderr
						cmd.Run()
					} else {
						cmd := exec.Command("./adb.exe", "shell", "pm", "disable-user", "--user", "0", app)
						cmd.Dir = "src/adb-windows"
						cmd.Stdout = os.Stdout
						cmd.Stderr = os.Stderr
						cmd.Run()
					}

				} else if runtime.GOOS == "darwin" {
					if action == "U" {
						cmd := exec.Command("./adb", "uninstall", app)
						cmd.Dir = "src/adb-mac"
						cmd.Stdout = os.Stdout
						cmd.Stderr = os.Stderr
						cmd.Run()
					} else {
						cmd := exec.Command("./adb", "shell", "pm", "disable-user", "--user", "0", app)
						cmd.Dir = "src/adb-mac"
						cmd.Stdout = os.Stdout
						cmd.Stderr = os.Stderr
						cmd.Run()
					}
				}
			}
		}

	}

}
