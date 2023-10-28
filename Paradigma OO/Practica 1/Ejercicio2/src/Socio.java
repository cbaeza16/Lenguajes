import java.util.ArrayList;
import java.util.Date;
import java.util.List;
class Socio {
    private int numeroSocio;
    private String nombre;
    private String direccion;
    private List<Prestamo> prestamos;

    public Socio(int numeroSocio, String nombre, String direccion) {
        this.numeroSocio = numeroSocio;
        this.nombre = nombre;
        this.direccion = direccion;
        this.prestamos = new ArrayList<>();
    }

    public int getNumeroSocio() {
        return numeroSocio;
    }

    public String getNombre() {
        return nombre;
    }

    //número de libros que un socio tiene prestados
    public int getNumPrestamos() {
        return prestamos.size();
    }

    public void realizarPrestamo(Libro libro, Date fechaPrestamo) {
        if (libro.isDisponible()) {
            prestamos.add(new Prestamo(libro, fechaPrestamo, this));
            libro.setDisponible(false);
        } else {
            System.out.println("El libro no está disponible para préstamo.");
        }
    }

    public List<Prestamo> getPrestamos() {
        return prestamos;
    }
}