new Vue({
  el: '#app',
  data: {
    searchTerm: '',
    columns: ["Date", "From", "To", "Subject"], // Reemplaza con los nombres de tus columnas
    data: [], // Reemplaza con los datos obtenidos de tu API
    error:'',
    specificData:'',
    boolData: true,
    selectedOption: '_all',
    sliderValue: 50,
    currentPage: 1,
    itemsPerPage:20
  },
  computed: {
    paginatedData() {
      const startIndex = (this.currentPage - 1) * this.itemsPerPage;
      const endIndex = startIndex + this.itemsPerPage;
      return this.data.slice(startIndex, endIndex);
    },
    totalPages() {
      return Math.ceil(this.data.length / this.itemsPerPage);
    },
  },
  methods: {
    setSpecificData(source){
      this.specificData = "Subject: "+ source.Subject + "\n" +
                          "From: "+ source.From + "\n" +         
                          "To: "+ source.To + "\n" +
                          "Date: "+ source.Date + "\n" +
                          "File-type: "+ source.File + "\n\n" +
                          source.Message
    },
    async search() {
      console.log("search...")
      this.boolData = false
      this.data = ""
      this.specificData = ""
      const inputData = {
        "search_type": "match",
        "query": {
          "term": this.searchTerm,
          "field": this.selectedOption
        },
        "sort_fields": ["-@timestamp"],
        "from": 0,
        "max_results": Number(this.sliderValue),
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
        // Recibe la información y lo almacena en "data" si tiene contenido, si no, se envia un mensaje de error        this.data = data.hits.hits
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
    },sortByColumn(column) {
      this.data = this.data.sort((a, b) => {
        const valueA = a._source[column];
        const valueB = b._source[column];
    
        // Verificar si la columna es una fecha
        if (column === "Date") {
          const dateA = new Date(valueA);
          const dateB = new Date(valueB);
          return dateA - dateB;
        } else {
          // Comparar los valores como texto
          return valueA.localeCompare(valueB);
        }
      });
    },
    formatDate(date) {
      return moment(date, "ddd, DD MMM YYYY HH:mm:ss Z (ZZ)").format("DD/MM/YYYY");
    },
    changePage(value) {
      if(value === "+"){
        if(this.currentPage < this.totalPages){
          this.currentPage = this.currentPage + 1;
        }
      }else{
        if(this.currentPage > 1){
        this.currentPage = this.currentPage - 1;
        }
      }
    },
  }
});