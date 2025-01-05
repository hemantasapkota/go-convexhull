go-convexhull
=============

Implementation of Graham Scan in GO with visualization

![SS](ss.png)

How to Use
=========
* Use mouse to add points on the screen. The hull is computed everytime a point is added.
* Press 'C' to clear the points.

**Build and Run**
```
git clone https://github.com/hemantasapkota/go-convexhull
cd go-convexhull && go build && ./go-convexhull
```

Changelog
=========
### 2025-01-05 - Migration to Raylib
* Replaced OpenGL/GLFW dependencies with Raylib for graphics rendering
* This change addresses the deprecation of OpenGL 3.0 in MacOS
* Improves cross-platform compatibility
* Simplifies graphics code while maintaining the same functionality
