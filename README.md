# OT
Script CLI para realizar una serie de optimizaciones a la tablet HTG 1078 con Soc RK 3066

Optimización de Tablet [OT]
Este pequeño proyecto tiene la finalidad de mejorar el rendimiento de las tablet HTG 1078 Gob con Soc rk3066 
debido a una version de android casi que obsoleta y la nesesidad de tener estas tablets en funcionamiento

Caracteristicas
1. Es una herramienta 100% CLI
2. Se basa en la herramienta ADB
3. Esta disponible para los 3 OS mas usados en ARM/x64

Requisitos

1. El ADB esta integrado por lo que no es nesesario descargarlo


CONFIGURACIÓN

el ejecutable esta acompañado de 2 archivos

1.install.text

    este se encarga de todas las apps se instalaran, al ingresar vera esta configuración

    app true I  -> indica que la app se instalara
    app false I -> la app no se instalara

    el nombre app es exactamente de el instalador apk que esta en la carpeta apks

    si quieres agregar nuevas apps junto con el script debes colocar el instalador en la carpeta apks, copiar el nombre
    pegarlo en install sin .apk acompañado de true e I

      PARA EL CORRECTO FUNCIONAMIENTO SOLO DEBE HABER UNA CONFIGURACIÓN POR LINEA



2.uninstall.txt

    similar a lo anterior tiene la estructura

    Paquete true U/D 

    paquete -> el servicio con el que se trabaja
    true -> si se tomara en cuenta o no 
    U/D -> la acción que se llevara a cabo, U - Uninstall , D - Disable

    TENGA EN CUENTA QUE HAY PAQUETES QUE NO SE PUEDEN ELIMINAR TOTALMENTE O DESACTIVAR PARA EL CORRECTO FUNCIONAMIENTO DE EL OS

    PARA EL CORRECTO FUNCIONAMIENTO SOLO DEBE HABER UNA CONFIGURACIÓN POR LINEA
El ejecutable se acompaña de 2 carpetas

  1.**apks** : donde estan todos los apk
  2.**src** : donde se guardan 3 subcarpetas
    1.**adb-linux**
    2.**adb-windows**
    3.**adb-mac**

    cada una de ellas tiene en su interior el adb

Uso


[Pasos previos]
1. tener la tablet formateada
    Ajustes -> Copia de Seguridad -> Restablecer datos de fábrica -> Restablecer Tablet (Borrar Targeta SD activada)

2. tener la tablet en modo de desarrollador y la depuracion USB activada
    Ajustes -> Información de la tablet -> (Presionar Número de compilación 7 veces)


3. la tablet solo se puede conectar una ve esten cumplidos los 2 anteriores

4. descargar el .zip
5. descomprimirlo en un lugar accesible
6. ejecutar el programa correspondiente a su OS / Arquitectura

[Windows]
    En windows el uso de la herramienta esta condicionado al uso de PowerShell para funcionar

    se debe abrir el archivo correspondiente a windows en su respectiva arquitectura y ejecutarlo
    si tiene problemas para funcionar se debe correr como admin

    ./OT_windows_x64.exe    
    ./OT_windows_arm.exe

[Mac/Linux]

    Se debe seleccionar el ejecutable correspondiente a tu OS y arquitectura

    ./OT_linux_x64/arm
    ./OT_mac_intel/arm

    generalmente en mac es probable que se deba ejecutar como admin para que funcione correctamente

este proyecto esta publicado bajo la licencia MIT consulte el archivo LICENSE
