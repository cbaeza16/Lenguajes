import java.util.ArrayList;
import java.util.List;
import java.util.Date;
import java.util.stream.Collectors;

public class Biblioteca {
    private static List<Socio> socios = new ArrayList<>();

    public static void main(String[] args) {

        //Crear socios
        Socio socio1 = new Socio(01, "Luna", "Aaljuela");
        Socio socio2 = new Socio(02, "Luis", "Grecia");
        Socio socio3 = new Socio(03, "Mario", "San Jose");
        Socio socio4 = new Socio(04, "Miriam", "Heredia");

        Libro libro1 = new Libro(11, "Libro1", "Autor1");
        Libro libro2 = new Libro(12, "Libro2", "Autor2");
        Libro libro3 = new Libro(13, "Libro3", "Autor3");
        Libro libro4 = new Libro(14, "Libro4", "Autor4");
        Libro libro5 = new Libro(15, "Libro5", "Autor5");
        Libro libro6 = new Libro(16, "Libro6", "Autor6");
        Libro libro7 = new Libro(17, "Libro7", "Autor7");

        //Registrar socios
        socios.add(socio1);
        socios.add(socio2);
        socios.add(socio3);
        socios.add(socio4);

        //Prestamos
        socio1.realizarPrestamo(libro1, new Date());
        socio1.realizarPrestamo(libro2, new Date());
        socio1.realizarPrestamo(libro3, new Date());
        socio1.realizarPrestamo(libro4, new Date());
        socio2.realizarPrestamo(libro5, new Date());

        //Mostrar el estado de los libros
        System.out.println("Libros:");
        System.out.println("Libro 1: Disponible: " + libro1.isDisponible());
        System.out.println("Libro 2: Disponible: " + libro2.isDisponible());
        System.out.println("Libro 3: Disponible: " + libro3.isDisponible());
        System.out.println("Libro 4: Disponible: " + libro4.isDisponible());
        System.out.println("Libro 5: Disponible: " + libro5.isDisponible());
        System.out.println("Libro 6: Disponible: " + libro6.isDisponible());
        System.out.println("Libro 7: Disponible: " + libro7.isDisponible());

        //Mostrar socios con más de 3 libros prestados
        List<Socio> sociosConMasDe3Libros = getSociosMas3Libros();
        System.out.println("Socios con más de 3 libros prestados:");
        for (Socio socio : sociosConMasDe3Libros) {
            System.out.println("Socio: " + socio.getNombre() + ", Libros prestados: " + socio.getNumPrestamos());
        }
    }

    //Obtener la lista de socios con más de 3 libros prestados
    public static List<Socio> getSociosMas3Libros() {
        return socios.stream()
                .filter(socio -> socio.getNumPrestamos() > 3)
                .collect(Collectors.toList());
    }
}