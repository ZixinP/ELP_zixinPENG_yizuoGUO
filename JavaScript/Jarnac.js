//https://github.com/sfrenot/javascript/tree/master/projet2

function createtable(container) {

    var tableContainer = document.getElementById(container);
    var table = document.createElement("table");
    table.style.borderCollapse = "collapse";
    table.style.width = "300px";
    table.style.height = "300px";
  
    for (var row = 1; row < 10; row++) {
      var tr = document.createElement("tr");
  
      for (var col = 1; col < 10; col++) {
        var td = document.createElement("td");
        td.style.border = "1px solid #000";
        td.style.width = "30px";
        td.style.height = "30px";
        //td.style.textAlign = "center";
        //td.style.verticalAlign = "middle";
        //td.appendChild(document.createTextNode(row * col));
        tr.appendChild(td);
      }
      table.appendChild(tr);
    }
    tableContainer.appendChild(table);
  
  }
  window.onload = function() {
    createtable("TableContainer");
    createtable("TableContainer2");
  };