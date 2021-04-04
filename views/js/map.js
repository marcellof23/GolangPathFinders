window.markers = [];
window.polylines = [];

let activeMarkers = [];

function initMap() {
	const bandung = { lat: -6.9175, lng: 107.6191 };

	window.map = new google.maps.Map(document.getElementById("map"), {
		zoom: 17,
		center: bandung,
	});

	window.map.addListener("click", (e) => {
		placeMarkerAndPanTo({ lat: e.latLng.lat(), lng: e.latLng.lng() });
	});
}

const graphForm = document.getElementById("graphForm");
const inpFile = document.getElementById("inpFile");

graphForm.addEventListener("submit", e =>{
    e.preventDefault();
    const reader = new FileReader();
    reader.onload = () => {

        for(let i=0;i<window.markers.length;i++){
            window.markers[i].setMap(null);
        }
        window.markers = [];

        const lines = reader.result.split('\n').map((line)=>{
            return line.split(',');
        });
        console.log(lines);
        console.log(parseInt(lines[0][0]));
        for(let i=1;i<1+parseInt(lines[0][0]);i++){
            placeMarkerAndPanTo({lat:parseFloat(lines[i][1]),lng:parseFloat(lines[i][2])});
        }
        for(let i=1+parseInt(lines[0][0]);i<1+2*parseInt(lines[0][0]);i++){
            console.log(lines[i][0]);
            let row = lines[i][0].split(" ");
            console.log(row);
            for(let j=0;j<parseInt(lines[0][0]);j++){
                if(parseInt(row[j])===parseInt("1")){
                    console.log(row[j] , "===" , "1" , "is " , String(row[j][0])==String("1"));
                    console.log(typeof row[j], typeof "1");
                    addLines({lat: markers[i-1-parseInt(lines[0][0])].position.lat(),lng:markers[i-1-parseInt(lines[0][0])].position.lng()},{lat: markers[j].position.lat(),lng: markers[j].position.lng()});
                }
            }
        }
    }
    reader.readAsText(inpFile.files[0]);
});

const placeMarkerAndPanTo = (coords) => {
    let marker = new google.maps.Marker({
      position: coords,
      map: window.map,
      icon: "http://maps.google.com/mapfiles/ms/icons/red-dot.png"
    });
    marker.addListener('click', (e)=>{
        console.log(marker.position.lat());
        console.log(marker.position.lng());
    })
    marker.addListener('contextmenu', (e)=>{
        marker.setIcon("http://maps.google.com/mapfiles/ms/icons/blue-dot.png");
    })

    window.markers.push(marker);
    window.map.panTo(coords);
}

const clearMarker = ()=>{
    for(let i=0;i<window.markers.length;i++){
        window.markers[i].setMap(null);
    }
    for(let i=0;i<window.polylines.length;i++){
        window.polylines[i].setMap(null);
    }

    window.markers = [];
    window.polylines = [];
}

const clearMarkers = document.getElementById("clearNodes");
clearMarkers.addEventListener('click',clearMarker);

const addLines = (sourceCoords,destinationCoords) => {
    let polyline = new google.maps.Polyline({
        path: [sourceCoords,destinationCoords],
        geodesic: true,
        strokeColor: "#FF0000",
        strokeOpacity: 1.0,
        strokeWeight: 2,
    })
    polyline.setMap(window.map);
    window.polylines.push(polyline);
}
