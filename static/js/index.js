function submitSearch(){
    $.ajax({
        url:"/search",
        method:"POST",
        data: $("#search-input").serialize(),
        success : function(rawData) {
            var jsonData = JSON.parse(rawData);
            if (!jsonData){
                return;
            }

            var searchResults = $("#search-results");
            searchResults.empty();

            jsonData.forEach(function(result){
               var row = $("<tr><td>" + result.Title +"</td><td>" + result.Author +"</td><td>" + result.Year +"</td><td>" + result.ID +"</td></tr>")
                searchResults.append(row);
            });
        }
    });
    return false;
}