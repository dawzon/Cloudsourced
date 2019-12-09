const base_url = 'http://' + window.location.hostname + ':8080/api'

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

export async function uploadFile() {
    //TODO
}

export async function deleteFile() {
    //TODO
}

export async function submitJob(name, description, platform, minram, files, script) {
    
    var request = {
        "name": name,
        "description": description,
        "platform": platform,
        "minram": minram,
        "files": files,
        "script": script,
    }

    return fetch(base_url + 'submit_job', {method: "POST", body: JSON.stringify(request)})
}

export async function cancelJob() {
    //TODO
}

export async function getActiveNodes(){

    const response = await fetch(base_url + "/active_nodes")
    return response.json()
}