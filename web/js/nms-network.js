var nodes = null;
var edges = null;
var network = null;
console.log("===========")
var DIR = 'img/';
var EDGE_LENGTH_MAIN = 150;
var EDGE_LENGTH_SUB = 50;

// Called when the Visualization API is loaded.
function draw() {
    // Create a data table with nodes.
    nodes = [];

    // Create a data table with links.
    edges = [];

    nodes.push({ id: 1, label: 'kube-master', image: DIR + 'ubuntu.png', shape: 'image' });
    nodes.push({ id: 2, label: 'Office', image: DIR + 'linux.png', shape: 'image' });
    nodes.push({ id: 3, label: 'kuber-1\n(10.100.1.1)', image: DIR + 'pod.png', shape: 'image' });
    nodes.push({ id: 4, label: 'kuber-2', image: DIR + 'pod.png', shape: 'image' });
    edges.push({ from: 1, to: 2 })
    edges.push({ from: 1, to: 3 })
    edges.push({ from: 1, to: 4 })
        // create a network
    var container = document.getElementById('network');
    var data = {
        nodes: nodes,
        edges: edges
    };
    $.get("http://localhost:8000/api/network/ns", function(rs) {
        var options = {};
        console.log(rs.Data)
        for (let index = 0; index < rs.Data.nodes.length; index++) {
            const o = rs.Data.nodes[index];
            o.shape = "image";
            switch (o.type) {
                case "pod":
                    o.image = DIR + "pod.png";
                    break;
                default:
                    o.image = DIR + "linux.png";
                    break;
            }
            o.label = o.label + "\n" + o.ip;
            rs.Data.nodes[index] = o
        }
        data = rs.Data;
        network = new vis.Network(container, data, options);
    })
}
$(function() { draw() })