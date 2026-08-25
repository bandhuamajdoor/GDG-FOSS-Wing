const fs = require('fs');

function _chk(r) {
    if (r.length !== 3) return false;
    if (r[1] !== 'IN' && r[1] !== 'OUT') return false;
    return true;
}

function processLog(data) {
    data.sort((a, b) => parseInt(a[2]) - parseInt(b[2]));
    
    let st = {};
    let res = {};
    
    for (let i = 0; i < data.length; i++) {
        let p = data[i][0];
        let ax = data[i][1];
        let t = data[i][2];
        
        if (!_chk([p, ax, t])) continue;
        
        if (ax === 'IN') {
            if(st[p]==undefined){    // fix 1: we delete intime on exit, hence check if an intime already exists
            st[p] = parseInt(t);
        }
        } else if (ax === 'OUT') {
            if (st[p] !== undefined) {
                let d = parseInt(t) - st[p];
                if (d < 0) d = 0;

                if(res[p]==undefined) res[p]=0;   //fix 2.0:initialise total time 
                // Calculate active duration
                res[p]+=d;                        //fix 2.0: add duration incase of multiple shifts,dont update
                
                let temp = st[p]; // <-- The 3-year-old temp variable
                delete st[p];         
            }
        }
    }
    return res;
}

if (process.argv.length < 3) {
    console.log("Usage: node process_logs.js <logfile>");
    process.exit(1);
}

const raw = fs.readFileSync(process.argv[2], 'utf8')
    .split('\n')
    .filter(line => line.trim() !== '')
    .map(line => line.trim().split(/\s+/));

const metrics = processLog(raw);
Object.keys(metrics).sort().forEach(user => {
    console.log(`${user}: ${metrics[user]} mins`);
});