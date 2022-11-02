# MyFyne

## Prerequisites
for me in ArtixLinux

> sudo pacman -S go xorg-server-devel libxcursor libxrandr libxinerama libxi


## Start

when we start app

> go mod init appname

> go get fyne.io/fyne/v2

when we finish app run

> go tidy


## Helloworld



> w.ShowAndRun

> w.Show() && a.Run()

不要运行app多次，仅能运行一次




## Markdown

### 这里Import库要注意

'''go
import (
"fyne.io/fyne/v2"
"fyne.io/fyne/v2/app"
"fyne.io/fyne/v2/container"
"fyne.io/fyne/v2/widget"
)
'''

### 中文乱码解决

我拷贝了window的ttf到linux下

/usr/share/fonts/

采用了他人写好点字体查询函数将其设置成fyne的字体

## GoldWatch

