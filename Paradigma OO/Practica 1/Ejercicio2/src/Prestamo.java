import java.util.Date;
class Prestamo {
    private Libro libro;
    private Date fechaPrestamo;
    private Socio socio;

    public Prestamo(Libro libro, Date fechaPrestamo, Socio socio) {
        this.libro = libro;
        this.fechaPrestamo = fechaPrestamo;
        this.socio = socio;
    }

    public Libro getLibro() {
        return libro;
    }

    public Date getFechaPrestamo() {
        return fechaPrestamo;
    }

    public Socio getSocio() {
        return socio;
    }
}