public class Main {
    public static void main(String[] args) {
        Documento documento = new Documento();

        //Hoja 1
        Hoja hoja1 = new Hoja();
        Texto texto = new Texto(200, 20, "Prueba 1");
        Circulo circulo = new Circulo(30, 30, 5);

        hoja1.agregarObjeto(texto);
        hoja1.agregarObjeto(circulo);

        //Hoja 2
        Hoja hoja2 = new Hoja();
        Rectangulo rectangulo = new Rectangulo(40, 40, 20, 10);
        Cuadrado cuadrado = new Cuadrado(10,10,20);

        hoja2.agregarObjeto(rectangulo);
        hoja2.agregarObjeto(cuadrado);

        //Agregar hojas al documento
        documento.agregarHoja(hoja1);
        documento.agregarHoja(hoja2);

        documento.imprimirDocumento();
    }
}
