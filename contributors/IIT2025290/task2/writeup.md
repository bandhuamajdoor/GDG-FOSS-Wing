# What exactly was wrong, and why did it only show up on some inputs and not others? Name the specific line/condition.
One problem was that if koi IN karke OUT se pehle fir se IN karle to purana vala IN overwrite ho jata tha and hence pichla wala time record nahi hota tha
Fix - if (st[p] === undefined)
            {
                st[p] = parseInt(t); 
            }
. Also, another problem was that if koi user agar ek baar out karke chale jaaye i.e already exist karta hai to uske existing time mea add ho jayega new time if vo firse in and out karta hai. Ayr, agar pehli baar out kar rha hai to use apne aap 0 maankar usme add ho jayega.
Fix - res[p] = (res[p] || 0) + d;

# Paste your output for tricky_input.txt. Does your fix handle this one correctly? If yes, why is it robust; if no, what class of input still breaks it?
Output: 
Dave: 40 mins
Eve: 60 mins
Yes, my fix handles this one correctly but mera code phas jayega agar Koi IN karta hai but OUT nahi karta hai, to uska add hi nahi hoga time. Also, agar aisa input aa jata hai jiski line mea 3rd column mea number ki jagah kuch text vgrh diya hua hai to NaN throw karega. Baaki, my code handles all inputs.