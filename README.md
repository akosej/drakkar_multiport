
````markdown
# DRAKKAR- Proxy Multiport | localhost 🔀 WSL Port

Un pequeño programa en Go que permite redirigir puertos desde Windows hacia servicios corriendo en WSL (Windows Subsystem for Linux). Ideal para exponer servicios como servidores web o APIs desarrolladas dentro de WSL hacia la red local o a otras aplicaciones en Windows.

---

## 🚀 Características

- Redirección de puertos TCP desde Windows a WSL.
- IP configurable (por defecto `localhost`).
- Puede ejecutarse como herramienta CLI desde cualquier terminal.
- Código 100% en Go, sin dependencias externas ni uso de PowerShell.

---

## 🛠️ Instalación

### 1. Clona este repositorio

```bash
git clone https://github.com/akosej/drakkar_multiport.git
cd drakkar_multiport
````

### 2. Compila el ejecutable

#### En Windows:

```bash
go build -o drakkar.exe
```

#### En WSL/Linux para Windows:

```bash
GOOS=windows GOARCH=amd64 go build -o drakkar.exe
```

### 3. (Opcional) Agrega al PATH

Para usarlo como comando global:

* Copia `drakkar.exe` a una carpeta como `C:\Tools`
* Agrega `C:\Tools` al PATH de tu sistema

---

## ⚙️ Uso

### Sintaxis:

```bash
drakkar [puerto1] [puerto2] ...
```

Cada puerto especificado será redirigido desde `0.0.0.0:<puerto>` hacia `localhost:<puerto>` (dentro de WSL).

### Ejemplo:

```bash
drakkar 8080 3000 9000
```

Esto redirige:

* `localhost:8080` ⟶ `WSL:8080`
* `localhost:3000` ⟶ `WSL:3000`
* `localhost:9000` ⟶ `WSL:9000`

---

## 🧪 Requisitos

* Go 1.18 o superior
* Tener configurado WSL y los servicios corriendo en los puertos deseados dentro de WSL

---

## 📦 Distribución

Puedes incluir este ejecutable en tus scripts de desarrollo o integrarlo como herramienta auxiliar para exponer servicios fácilmente desde WSL.

---

## 📄 Licencia

MIT © \[Akos EJ]
