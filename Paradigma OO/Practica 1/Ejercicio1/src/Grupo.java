import java.util.ArrayList;
import java.util.List;

public class Grupo extends ObjetoRepresentable {
    private List<ObjetoRepresentable> objetosAgrupados;

    public Grupo(int x, int y) {
        super(x, y);
        objetosAgrupados = new ArrayList<>();
    }

    public void agregarObjeto(ObjetoRepresentable objeto) {
        objetosAgrupados.add(objeto);
    }

    public void eliminarObjeto(ObjetoRepresentable objeto) {
        objetosAgrupados.remove(objeto);
    }

    public void dibujar() {
        System.out.println("Dibujando grupo en (" + getX() + ", " + getY() + ")");
        for (ObjetoRepresentable objeto : objetosAgrupados) {
            objeto.dibujar();
        }
    }
}




