public class Circulo extends ObjetoGeometrico {
    private double radio;

    public Circulo(int x, int y, double radio) {
        super(x, y);
        this.radio = radio;
    }

    public double calcularArea() {
        return Math.PI * radio * radio;
    }

    public void dibujar() {
        System.out.println("Agregando cirtculo en (" + getX() + "| " + getY()+ ")" );
    }
}

