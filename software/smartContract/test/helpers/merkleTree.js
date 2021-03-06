const { keccak256, bufferToHex } = require('ethereumjs-util');

class MerkleTree {
  constructor (elements, hashFunction = keccak256) {
    this.hashFunction = hashFunction;
    this.layers = this.getLayers(elements.map(hashFunction));
  }

  getLayers (elements) {
    let emptyLeveled = this.hashFunction('');
    if (elements.length % 2) {
      elements.push(emptyLeveled);
    }

    const tree = [elements];
    const maxLevel = 31;
    for (let level = 1; level <= maxLevel; level++) {
      const current = [];
      for (let i = 0; i < tree[level - 1].length / 2; i++) {
        const a = tree[level - 1][i*2];
        const b = tree[level - 1][i*2 + 1];
        const hash = this.hashFunction(Buffer.concat([a, b]));
        current.push(hash);
      }

      if (current.length % 2 && level < maxLevel) {
        current.push(emptyLeveled);
      }
      emptyLeveled = this.hashFunction(Buffer.concat([emptyLeveled, emptyLeveled]));

      tree.push(current);
    }
    return tree;
  }

  getRoot () {
    return this.layers[this.layers.length - 1][0];
  }

  getHexRoot () {
    return bufferToHex(this.getRoot());
  }

  getProof (idx) {
    if (idx === -1) {
      throw new Error('Element does not exist in Merkle tree');
    }
    return this.layers.reduce((proof, layer) => {
      const pairElement = this.getPairElement(idx, layer);
      if (pairElement) {
        proof.push(pairElement);
      }

      idx = Math.floor(idx / 2);
      return proof;
    }, []);
  }

  getHexProof (idx) {
    const result = this.getProof(idx);
    return this.bufArrToHexArr(result);
  }

  getPairElement (idx, layer) {
    const pairIdx = idx % 2 === 0 ? idx + 1 : idx - 1;
    if (pairIdx < layer.length) {
      return layer[pairIdx];
    } else {
      return null;
    }
  }

  bufIndexOf (el, arr) {
    let hash;

    // Convert element to 32 byte hash if it is not one already
    if (el.length !== 32 || !Buffer.isBuffer(el)) {
      hash = this.hashFunction(el);
    } else {
      hash = el;
    }

    for (let i = 0; i < arr.length; i++) {
      if (hash.equals(arr[i])) {
        return i;
      }
    }

    return -1;
  }

  bufArrToHexArr (arr) {
    if (arr.some(el => !Buffer.isBuffer(el))) {
      throw new Error('Array is not an array of buffers');
    }

    return arr.map(el => '0x' + el.toString('hex'));
  }
}

function keccak160 (element) {
  return keccak256(element).slice(12,32);
}

module.exports = {
  MerkleTree,
  keccak160
};
