public class Cuadrado extends ObjetoGeometrico {
    private double lado;

    public Cuadrado(int x, int y, double lado) {
        super(x, y);
        this.lado = lado;
    }

    public double calcularArea() {
        return lado * lado;
    }

    public void dibujar() {
        System.out.println("Agregando cuadrado en (" + getX() + "| " + getY()+ ")" );
    }
}
