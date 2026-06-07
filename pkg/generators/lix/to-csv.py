# pip install pandas pyreadstat


from   pathlib import Path
import pandas as pd
import pyreadstat


def stackTrace(e):
    print(e)




def main():
    srcPth   = Path("tmp-bruttostichprobe.dta")
    mixinPth = srcPth.with_name("mixin.csv")
    dstPth   = srcPth.with_name("invitation.csv")

    try:

        df1, meta = pyreadstat.read_dta(str(srcPth), encoding="utf-8")

        if False:
            # print(meta)
            # print(df.head)
            df1 = pd.DataFrame(df1)
            print(df1.describe)
            # 1.443.172 Mio rows

        # renaming
        df1 = df1.rename(columns={"email_firma": "email"})

        df2 = pd.read_csv(mixinPth, sep="\t", dtype=str, encoding="utf-8", keep_default_na=False)
        if "userid" not in df2.columns:
            raise ValueError("missing userid column in mixin.csv")
        if "link" not in df2.columns:
            raise ValueError("missing link column in mixin.csv")


        # limiting both inputs
        cntLimit = 5000
        df1      = df1.head(cntLimit).copy()
        df2      = df2.head(cntLimit).copy()
        ln1      = len(df1)
        ln2      = len(df2)
        if ln1 != ln2:
            raise(f"row count mismatch: dta={ln1}, mixin={ln2}")


        # inserting mixin values as second and third CSV columns
        df1.insert(1, "userid", df2["userid"].to_list())
        df1.insert(2, "link",   df2["link"  ].to_list())

        df1.to_csv(dstPth, index=False, encoding="utf-8", sep=";")


    except Exception as exc:
        stackTrace(exc)
        print("convertStataToCsv")


if __name__ == "__main__":
    main()
