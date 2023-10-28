
public abstract class ObjetoGeometrico extends ObjetoRepresentable {
    public ObjetoGeometrico(int x, int y) {
        super(x, y);
    }

    public abstract double calcularArea();
}