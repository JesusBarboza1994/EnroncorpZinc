new Vue({
  el: '#app',
  data: {
    searchTerm: '',
    columns: ["Date", "From", "To", "Subject", "User", "File"], // Reemplaza con los nombres de tus columnas
    data: [], // Reemplaza con los datos obtenidos de tu API
    error:'',
    specificData:''
    
  },
  computed: {
    sortedData() {
      // Aplica el filtro y ordena los datos según la columna seleccionada
      return this.data.filter(item => {
        return Object.values(item).some(value =>
          String(value).toLowerCase().includes(this.searchTerm.toLowerCase())
        );
      }).sort((a, b) => {
        // Cambia 'field1' por la columna que deseas ordenar inicialmente
        return String(a['field1']).localeCompare(String(b['field1']));
      });
    }
  },
  methods: {
    setSpecificData(source){
      this.specificData = source.Message

    },
    async search() {
      console.log("search...")
      console.log(this.searchTerm)
      this.data = ""
      const inputData = {
        "search_type": "match",
        "query": {
          "term": this.searchTerm,
          "field": "_all"
        },
        "sort_fields": ["-@timestamp"],
        "from": 0,
        "max_results": 20,
        "_source": []
      };
  
      try {
        const response = await fetch('http://localhost:8000/search', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(inputData)
        });
        const data = await response.json();
        console.log(data)
        // Recibe la información y lo almacena en "data" si tiene contenido, sino, se envia un mensaje de error        this.data = data.hits.hits
        if (data.hits.hits.length > 0) {
          this.data = data.hits.hits
          this.error = ""
        } else {
          this.error = "No se encontraron resultados"
        }
      } catch (error) {
        // Recibe el error
        this.error = error
        console.error(error);
      }
    },
    formatDate(date) {
      return moment(date, "ddd, DD MMM YYYY HH:mm:ss Z (ZZ)").format("DD/MM/YY");
    }
  }
});