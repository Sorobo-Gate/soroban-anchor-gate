import {
  Address,
  Contract,
  nativeToScVal,
  scValToNative,
  xdr,
} from "@stellar/stellar-sdk";

export interface EscrowDetails {
  payer: string;
  beneficiary: string;
  token: string;
  amount: bigint;
  profileHash: string;
  unlockTimestamp: bigint;
}

export class EscrowGateClient {
  private contract: Contract;

  constructor(public readonly contractId: string) {
    this.contract = new Contract(contractId);
  }

  /**
   * Encodes create_escrow invocation parameters for wallet signing.
   */
  public buildCreateEscrowTx(params: {
    payer: string;
    beneficiary: string;
    token: string;
    amount: bigint;
    profileHashHex: string;
    lockDurationSeconds: bigint;
  }): xdr.ScVal[] {
    const hashBytes = Buffer.from(params.profileHashHex, "hex");
    if (hashBytes.length !== 32) {
      throw new Error("profileHash must be exactly 32 bytes (64 hex characters)");
    }

    return [
      new Address(params.payer).toScVal(),
      new Address(params.beneficiary).toScVal(),
      new Address(params.token).toScVal(),
      nativeToScVal(params.amount, { type: "i128" }),
      xdr.ScVal.scvBytes(hashBytes),
      nativeToScVal(params.lockDurationSeconds, { type: "u64" }),
    ];
  }

  /**
   * Encodes release_to_anchor invocation parameters.
   */
  public buildReleaseToAnchorTx(params: {
    escrowId: bigint;
    caller: string;
    anchorDisbursementAddress: string;
  }): xdr.ScVal[] {
    return [
      nativeToScVal(params.escrowId, { type: "u64" }),
      new Address(params.caller).toScVal(),
      new Address(params.anchorDisbursementAddress).toScVal(),
    ];
  }

  /**
   * Encodes refund invocation parameters.
   */
  public buildRefundTx(escrowId: bigint): xdr.ScVal[] {
    return [nativeToScVal(escrowId, { type: "u64" })];
  }
}
