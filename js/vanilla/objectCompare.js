function getModifiedConfig(original, modified) {
    function compareObjects(orig, mod) {
        if (!orig || !mod || typeof orig !== typeof mod || 
            Array.isArray(orig) !== Array.isArray(mod)) {
            return mod;
        }

        if (Array.isArray(orig)) {
            return JSON.stringify(orig) !== JSON.stringify(mod) ? mod : undefined;
        }

        if (typeof orig !== 'object') {
            return orig !== mod ? mod : undefined;
        }

        const changes = {};
        const allKeys = new Set([...Object.keys(orig), ...Object.keys(mod)]);

        for (const key of allKeys) {
            if (!(key in orig)) {
                changes[key] = mod[key];
            } else if (!(key in mod)) {
                changes[key] = undefined;
            } else {
                const diff = compareObjects(orig[key], mod[key]);
                if (diff !== undefined) {
                    changes[key] = diff;
                }
            }
        }

        return Object.keys(changes).length > 0 ? changes : undefined;
    }

    return compareObjects(original, modified) || {};
}

// Example usage:
const original = {
    name: "John",
    age: 30,
    address: {
        street: "123 Main St",
        city: "Boston"
    },
    hobbies: ["reading", "gaming"],
    settings: {
        theme: "dark",
        notifications: {
            email: true,
            push: false
        }
    }
};

const modified = {
    name: "John",
    age: 31,
    address: {
        street: "456 Oak St",
        city: "Boston"
    },
    hobbies: ["reading", "gaming", "hiking"],
    settings: {
        theme: "light",
        notifications: {
            email: true,
            push: true
        }
    }
};

console.log(getModifiedConfig(original, modified));

/* Output:
{
    age: 31,
    address: {
        street: "456 Oak St"
    },
    hobbies: ["reading", "gaming", "hiking"],
    settings: {
        theme: "light",
        notifications: {
            push: true
        }
    }
}
*/
