let token      = null;
tableBody      = document.getElementById('table_01');
eksClusterList = document.getElementById('eks_cluster_list');
eksClusters    = [];
eksNodes       = [];
eksNamespaces  = [];
eksPods        = [];

// eksClusters.push
// tableBody.getElementsByTagName

function login() {
    fetch('http://localhost:8080/login', {
        method: 'POST'
    })
    .then(response => response.json())
    .then(data => {
        token = data.token;
    })
    .catch(error => {
        console.error('Login error:', error);
    });
}

window.addEventListener("DOMContentLoaded", (event) => {
    document.getElementById('btnGetClusters').addEventListener('click', function() {
        fetch('http://localhost:8080/list-clusters', {
            headers: {
                'Authorization': token
            }
        })
        .then(response => response.json())
        .then(data => {
            let tCount = 0;
            tableBody.innerHTML = '';
            caasClusters = data.clusters;
            console.log(caasClusters);
            if (eksClusters.length == 0 ) {
                caasClusters.forEach (function (cluster) {
                    tCount +=1 ;
                    if ( eksClusterList.hasOwnProperty(cluster) ) {
                        console.log("eksClusterList: " + eksClusterList[cluster]);
                    }

                    eksClusters.push(cluster);             // add cluster name to the Global list

                    // Populating dropdown list with entries
                    var listEntry = document.createElement('option');
                    listEntry.innerHTML = cluster;
                    eksClusterList.appendChild(listEntry);

                });
                eksClusterList = document.getElementById("eks_cluster_list");
                eksClusterList.options[0].innerHTML  = eksClusters.length + " clusters found";
            } else {
                console.log("Clusters found: " + eksClusters.length);
            }
        })
        .catch(error => {
            szErrorString = "";
            console.log('Error: ', error);
            outputBox = document.getElementById('output');

            if ( error.message.indexOf('Failed to fetch') !== -1 ) {
                szErrorString = '<table border=1><tr><td>ERROR</td>' +
                '<td> Failed to fetch EKS cluster information! <br>' +
                'Possible Cause: Breksit backend server is not running or not accessible.</td></tr></table>';
            }
            outputBox.innerHTML = szErrorString;
            if ( error.error ) { console.log('Error.error: ', error); }
            if ( error.code )  { console.log('Error.code: ',  code ); }
        });
    });
});

document.addEventListener('DOMContentLoaded', function() {
    tableBody       = document.getElementById('table_01');
    eksClusterList  = document.getElementById('eks_cluster_list');
});

// login();
