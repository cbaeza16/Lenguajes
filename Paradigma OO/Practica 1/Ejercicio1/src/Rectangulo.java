public class Rectangulo extends ObjetoGeometrico {
    private double ancho;
    private double alto;

    public Rectangulo(int x, int y, double ancho, double alto) {
        super(x, y);
        this.ancho = ancho;
        this.alto = alto;
    }

    public double calcularArea() {
        return ancho * alto;
    }

    public void dibujar() {
        System.out.println("Agregando rectangulo en (" + getX() + "| " + getY() + ")" );
    }
}

