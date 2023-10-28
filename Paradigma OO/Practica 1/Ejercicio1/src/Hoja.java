import java.util.ArrayList;
import java.util.List;

public class Hoja {
    private List<ObjetoRepresentable> objetosHoja;

    public Hoja() {
        objetosHoja = new ArrayList<>();
    }

    public void agregarObjeto(ObjetoRepresentable objeto) {
        objetosHoja.add(objeto);
    }

    public void eliminarObjeto(ObjetoRepresentable objeto) {
        objetosHoja.remove(objeto);
    }

    public void dibujarHoja() {
        System.out.println("Dibujando hoja:");
        for (ObjetoRepresentable objeto : objetosHoja) {
            objeto.dibujar();
        }
    }
}
