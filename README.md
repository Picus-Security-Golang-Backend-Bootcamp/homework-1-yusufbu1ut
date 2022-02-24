# Kitaplık App/ 1. Ödev 
Bu çalışmada os, fmt, strings, bufio(dosyadan veri okumak için kullanılan paket)  ve benim oluşturmuş olduğum helper paketleri ile uygulama gerçekleştirilmiştir.

## Helper Pakedi

Bu paket iki adet fonksiyon içermektedir. 

Bunlar;

- List(slice)

    Bu fonksiyon içerisinde kitapları içeren bir slice parametresi alarak slice içerisindeki verileri çıktılamaktadır.
- Search(string,slice)

    Bu fonksiyon içerisinde aranacak argumanı string olarak ve içerisinde aranacak kitaplığı slice olarak alır. String değeri ve slice değerleri küçük,büyük harf yapısına çevrilerek karşılaştırmalar yapılır. Benzerlik içeren veriler çıktılanmaktadır.

## main.go

Burada kitapları.txt dosyası okunarak kitap verileri alınır ve bir slice yapısı içerisine append işlemleri gerçekleştirilir, hata kontrolleri yapılır. 

Konsol üzerinden gelen argüman sayıları incelenir, ilk argüman incelenir ve argüman doğrultusunda helper üzerinden işlemler gerçekleştirilir. 

