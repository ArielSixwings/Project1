PHONY: clean

$(TARGET): MatPixel.a
	go build .

Matpixel.a: Matpixel.o
	ar r $@ $^

%.o: %.cpp
	g++ -O2 -o $@ -c $^

clean:
	rm -f *.o *.so *.a $(TARGET)
