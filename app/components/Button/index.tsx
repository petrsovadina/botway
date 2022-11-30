import { LoadingDots } from "@/components/LoadingDots";
import clsx from "clsx";
import { forwardRef } from "react";
import styles from "./Button.module.scss";

export const Button = forwardRef(function Button(
  { children, onClick, loading, disabled, className }: any,
  ref: any
) {
  return (
    <div className="mt-6 space-y-2 flex justify-center">
      <button
        className={clsx(
          "bg-blue-700 text-white hover:text-gray-200 transition button p-2",
          className
        )}
        ref={ref}
        onClick={onClick}
        disabled={loading || disabled}
      >
        {loading && <LoadingDots className={styles.loading} children />}
        <span>{children}</span>
      </button>
    </div>
  );
});
