public class Texto extends ObjetoRepresentable {
    private String texto;

    public Texto(int x, int y, String texto) {
        super(x, y);
        this.texto = texto;
    }

    public void dibujar() {
        System.out.println("Agregando texto en (" + getX() + "| " + getY() + "): " + texto);
    }
}