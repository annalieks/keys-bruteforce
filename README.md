# Keys Brute Force

1. Number of keys is `2^n` for an n-bit key
2. Generate a random key from an n-bits keys field. The algorithm generates a key by setting its bits to either 0 or 1, hence, the distribution is non-uniform (the probability of generating the key `k` is not the same for all `k` from the field). This way is better for demonstration purposes because the key will be large enough for brute force, but will less likely reach the max value.
3. Brute force a field of n-bits keys to find a match with a randomly generated one