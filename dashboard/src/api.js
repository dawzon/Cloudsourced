const base_url = 'http://' + window.location.hostname + ':8080/api'

export var alias

export function setAlias(a) {
    alias = a
}

export async function getRunningJobs() {

    const response = await fetch(base_url + '/running_jobs')
    return response.json()
}

export async function getQueuedJobs() {
    
    const response = await fetch(base_url + "/queued_jobs")
    return response.json()
}

export async function getFailedJobs() {
    const response = await fetch(base_url + "/failed_jobs")
    return response.json()
}

export async function getFinishedJobs() {
    const response = await fetch(base_url + "/finished_jobs")
    return response.json()
}

export async function getJobsByOwner(owner) {

    const response = await fetch(base_url + '/jobs_by_owner', {method: 'POST', body: JSON.stringify({"owner": owner})})
    return response.json()
}

// export async function uploadFile() {
//     //TODO
// }

// export async function deleteFile() {
//     //TODO
// }

export async function submitJob(name, description, platform, script, owner) {
    
    var request = {
        "name": name,
        "description": description,
        "platform": platform,
        "script": script,
        "owner": owner,
    }

    const response = await fetch(base_url + '/submit_job', {method: "POST", body: JSON.stringify(request)})
    return await response.json()
}

// export async function cancelJob() {
//     //TODO
// }

export async function getActiveNodes(){

    const response = await fetch(base_url + "/active_nodes")
    return response.json()
}

//TODO this is probably unnecessary
// export async function markAsActive(name) {

//     var request = {
//         "name": name
//     }
//     await fetch(base_url + "/connect", {method: "POST", body: JSON.stringify(request)})
// }

export async function getWork(alias) {

    var request = {
        "alias": alias
    }
    const response = await fetch(base_url + "/get_work", {method: "POST", body: JSON.stringify(request)})
    return await response.json()
}

export async function submitWork(alias, output) {

    //TODO
    var request = {
        "alias": alias,
        "output": output
    }
    const reponse = await fetch(base_url + "/submit_work", {method: "POST", body: JSON.stringify(request)})
}