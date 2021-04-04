window.markers = [];
window.polylines = [];

function initMap() {
	const bandung = { lat: -6.9175, lng: 107.6191 };

	window.map = new google.maps.Map(document.getElementById("map"), {
		zoom: 13,
		center: bandung,
	});

	window.map = new google.maps.Map(document.getElementById("map"), {
		zoom: 17,
		center: bandung,
	});

	window.map.addListener("click", (e) => {
		placeMarkerAndPanTo({ lat: e.latLng.lat(), lng: e.latLng.lng() });
	});
}

const placeMarkerAndPanTo = (coords) => {
	let marker = new google.maps.Marker({
		position: coords,
		map: window.map,
		color: "green",
	});
	marker.addListener("click", (e) => {
		console.log(marker.position.lat());
		console.log(marker.position.lng());
	});
	marker.addListener("contextmenu", (e) => {
		marker.setIcon("http://maps.google.com/mapfiles/ms/icons/blue-dot.png");
	});

	window.markers.push(marker);
	window.map.panTo(coords);
};

const clearMarker = () => {
	for (let i = 0; i < window.markers.length; i++) {
		window.markers[i].setMap(null);
	}
	for (let i = 0; i < window.polylines.length; i++) {
		window.polylines[i].setMap(null);
	}

	window.markers = [];
	window.polylines = [];
};

const clearMarkers = document.getElementById("clearNodes");
clearMarkers.addEventListener("click", clearMarker);
