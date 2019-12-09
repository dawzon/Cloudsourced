const base_url = 'http://' + window.location.hostname + ':8080/api'

export async function getRunningJobs() {

    const response = await fetch(base_url + '/running_jobs')
    return response.json()
}

export function getQueuedJobs() {
    //TODO
}

export function getFailedJobs() {
    //TODO
}

export async function getJobsByOwner(owner) {

    const response = await fetch(base_url + '/jobs_by_owner', {method: 'POST', body: JSON.stringify({"owner": owner})})
    return response.json()
}

export function uploadFile() {
    //TODO
}

export function deleteFile() {
    //TODO
}

export function submitJob() {
    //TODO
}

export function cancelJob() {
    //TODO
}