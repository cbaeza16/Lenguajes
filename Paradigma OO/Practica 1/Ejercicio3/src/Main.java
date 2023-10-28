public class Main {
    public static void main(String[] args) {
        // Obtén la instancia de la Agenda utilizando el Singleton
        Agenda agenda = Agenda.getInstance();

        // Utiliza la AgendaFactory para crear contactos y eventos
        AgendaFactory factory = new ContactoFactory();

        Contacto contactoPersonal = factory.crearContacto("Luna", "Miranda", "70854467", 22, "Personal");
        Contacto contactoEmpresarial = factory.crearContacto("Max", "Lopez", "708904467", 25, "Empresarial");

        factory = new EventoFactory();
        Evento eventoPersonal = factory.crearEvento("Cumpleannos", "18/09/23", "Personal");
        Evento eventoEmpresarial = factory.crearEvento("Reunión Zoom", "22/10/23", "Empresarial");

        // Agregar contactos y eventos a la agenda
        agenda.agregarContacto(contactoPersonal);
        agenda.agregarContacto(contactoEmpresarial);
        agenda.agregarEvento(eventoPersonal);
        agenda.agregarEvento(eventoEmpresarial);

        // Mostrar contactos y eventos
        agenda.mostrarContactos();
        agenda.mostrarEventos();
    }
}


//Diferencias entre un Eager Singleton y un Lazy Singleton

//En el Lazy Singleton la instancia de la clase singleton se crea solo cuando se requiere por primera vez.
//Ademas, la instancia no se crea durante el proceso de carga de clases.
//En el Eager Singleton a instancia de la clase singleton se crea con entusiasmo durante el
//proceso de carga de la clase, independientemente de si es necesaria o no.
//Ademas, garantiza que la instancia este disponible desde el inicio.

//Mejor a utilizar en la implementación del problema y por qué?

//Se va a utilizar el Lazy Singleton ya que es mas eficiente, hablando en recursos,
//ya que la instancia Agenda solo se creara cuando sea necesario.