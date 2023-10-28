import java.util.ArrayList;
import java.util.List;

public class Documento {
    private List<Hoja> hojas;

    public Documento() {
        hojas = new ArrayList<>();
    }

    public void agregarHoja(Hoja hoja) {
        hojas.add(hoja);
    }

    public void eliminarHoja(Hoja hoja) {
        hojas.remove(hoja);
    }

    public void imprimirDocumento() {
        System.out.println("Imprimiendo documento:");
        for (Hoja hoja : hojas) {
            hoja.dibujarHoja();
        }
    }
}
