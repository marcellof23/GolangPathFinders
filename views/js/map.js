let markers=[];

function initMap() {

    const bandung = { lat: -6.9175, lng: 107.6191 };

    window.map = new google.maps.Map(document.getElementById("map"), {
        zoom: 13,
        center: bandung,
    });

    window.map.addListener('click', (e)=>{
        placeMarkerAndPanTo({lat: e.latLng.lat(),lng: e.latLng.lng()});
    })
}

const placeMarkerAndPanTo = (coords) => {
    let marker = new google.maps.Marker({
      position: coords,
      map: window.map,
    });
    marker.addListener('click', (e)=>{
        console.log(marker.position);
        console.log(marker.position.lat());
        console.log(marker.position.lng());
    })
    marker.addListener('contextmenu', (e)=>{
        e.preventDefault();
        console.log("HAHAHAHAHAHA");
        console.log(marker.position.lat());
        console.log(marker.position.lng());
    })

    markers.push(marker);
    window.map.panTo(coords);
}

const clearMarker = ()=>{
    for(let i=0;i<markers.length;i++){
        markers[i].setMap(null);
    }
    markers = [];
}

const clearMarkers = document.getElementById("clearNodes");
clearMarkers.addEventListener('click',clearMarker);